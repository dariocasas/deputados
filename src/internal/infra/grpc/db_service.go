package grpc_service

import (
	"context"
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/infra/grpc/pb"
	"github.com/dariocasas/deputados3/src/internal/usecase/db_exists_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/drop_db_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/fetch_deputados_ids_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/populate_db_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/populate_index_usecase"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
)

type DbService struct {
	pb.UnimplementedDbServiceServer
	config                   *configs.Config
	dropDbUseCase            drop_db_usecase.DropDbUseCaseInterface
	dbExistsUseCase          db_exists_usecase.DbStatusUseCaseInterface
	fetchDeputadosIdsUseCase fetch_deputados_ids_usecase.FetchDeputadosIdsUseCaseInterface
	populateIndexUseCase     populate_index_usecase.PopulateIndexUseCaseInterface
	populateDbUseCase        populate_db_usecase.PopulateDbUseCaseInterface
}

func NewDbService(
	config *configs.Config,
	dropDbUseCase drop_db_usecase.DropDbUseCaseInterface,
	dbExistsUseCase db_exists_usecase.DbStatusUseCaseInterface,
	fetchDeputadosIdsUseCase fetch_deputados_ids_usecase.FetchDeputadosIdsUseCaseInterface,
	populateDbUseCase populate_db_usecase.PopulateDbUseCaseInterface,
	populateIndexUseCase populate_index_usecase.PopulateIndexUseCaseInterface,
) *DbService {
	return &DbService{
		config:                   config,
		dbExistsUseCase:          dbExistsUseCase,
		dropDbUseCase:            dropDbUseCase,
		fetchDeputadosIdsUseCase: fetchDeputadosIdsUseCase,
		populateDbUseCase:        populateDbUseCase,
		populateIndexUseCase:     populateIndexUseCase,
	}
}

func (s *DbService) DropDb(ctx context.Context, in *pb.DropDbRequest) (*pb.DropDbResponse, error) {

	log.Println("DbService.DropDb()")

	output := s.dropDbUseCase.Execute()
	if output.Error != nil {
		return nil, output.Error
	}
	return &pb.DropDbResponse{}, nil
}

func (s *DbService) DbStatus(ctx context.Context, in *pb.DbStatusRequest) (*pb.DbStatusResponse, error) {

	log.Println("DbService.PopulateIndex()")

	output := s.dbExistsUseCase.Execute()
	if output.Error != nil {
		return nil, output.Error
	}
	return &pb.DbStatusResponse{
		Status:     pb.DbStatusEnum(output.DbStatusEnum),
		IndexCount: output.IndexCount,
	}, nil
}

func (s *DbService) PopulateIndex(ctx context.Context, in *pb.PopulateIndexRequest) (*pb.PopulateIndexResponse, error) {

	log.Println("DbService.PopulateIndex()")

	output := s.fetchDeputadosIdsUseCase.Execute(
		fetch_deputados_ids_usecase.NewFetchDeputadosIdsInputDTO(
			s.config.HttpClientTimeOutMilliSecs,
			int(in.MaxItems),
		),
	)
	if output.Error != nil {
		return nil, output.Error
	}

	populateIndexOutputDTO := s.populateIndexUseCase.Execute(
		populate_index_usecase.NewPopulateIndexInputDTO(int(in.MaxItems),
			output.List,
		),
	)
	if populateIndexOutputDTO.Error != nil {
		return nil, populateIndexOutputDTO.Error
	}

	return &pb.PopulateIndexResponse{
		IndexCount: int32(len(output.List)),
	}, nil
}

func NewDeputadoResponse(
	id int, name string, error error,
	elapsedTime elapsed_time.ElapsedTimeInterface,
	partialTime populate_db_usecase.PartialTime,
) *pb.DeputadoResponse {
	var errorString string
	if error != nil {
		errorString = error.Error()
	}
	responsePartialTime := &pb.PartialTime{
		Dbread:   int32(partialTime.DbRead.Milliseconds()),
		Getapi:   int32(partialTime.GetApi.Milliseconds()),
		Getpage:  int32(partialTime.GetPage.Milliseconds()),
		Scanpage: int32(partialTime.ScanPage.Milliseconds()),
		Dbwrite:  int32(partialTime.DbWrite.Milliseconds()),
	}

	return &pb.DeputadoResponse{
		Id:           int32(id),
		Name:         name,
		Error:        error != nil,
		ErrorString:  errorString,
		Milliseconds: int32(elapsedTime.Elapsed().Milliseconds()),
		PartialTime:  responsePartialTime,
	}
}

func (s *DbService) CancelPopulateDb(context.Context, *pb.CancelPopulateDbRequest) (*pb.CancelPopulateDbResponse, error) {
	log.Println("DbService.CancelPopulateDb()")

	s.populateDbUseCase.Cancel()

	return nil, nil
}

func (s *DbService) PopulateDb(in *pb.PopulateDbRequest, out pb.DbService_PopulateDbServer) error {

	log.Println("DbService.PopulateDb()")

	populateDbOutputDTO := s.populateDbUseCase.Execute(

		populate_db_usecase.PopulateDbInputDTO{
			TimeOut:     int(in.Timeout),
			Concurrency: int(in.Concurrency),
			OnNewItem: func(
				id int,
				name string,
				error error,
				elapsedTime elapsed_time.ElapsedTimeInterface,
				partialTime populate_db_usecase.PartialTime,
			) {
				log.Println(partialTime)
				err := out.Send(NewDeputadoResponse(id, name, error, elapsedTime, partialTime))
				if err != nil {
					log.Printf("PopulateDb OnNewItem Send: %s", err.Error())
					s.populateDbUseCase.Cancel()
				}
			},
		},
	)

	if populateDbOutputDTO.Error != nil {
		return populateDbOutputDTO.Error
	}

	return nil
}
