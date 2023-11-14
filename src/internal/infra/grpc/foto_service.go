package grpc_service

import (
	"github.com/dariocasas/deputados3/src/internal/infra/grpc/pb"
	"github.com/dariocasas/deputados3/src/internal/usecase/get_fotos_usecase"
)

type FotosService struct {
	pb.UnimplementedFotosServiceServer
	fotosUseCase get_fotos_usecase.FotosUsecaseInterface
}

func NewFotosService(
	fotosUseCase get_fotos_usecase.FotosUsecaseInterface,
) *FotosService {
	return &FotosService{
		fotosUseCase: fotosUseCase,
	}
}

func NewFotoDTO(
	id int,
	fotoUrl string,
	fotoUrl2 string,
	nome string,
	partido string,
	uf string,
	url string,
	redes *pb.Redes,
) *pb.FotoResponse {
	return &pb.FotoResponse{
		Id:       int32(id),
		Nome:     nome,
		FotoUrl:  fotoUrl,
		FotoUrl2: fotoUrl2,
		Partido:  partido,
		Uf:       uf,
		Url:      url,
		Redes:    redes,
	}
}

func (s *FotosService) GetFotos(in *pb.GetFotosRequest, out pb.FotosService_GetFotosServer) error {

	FotosOutputDTO := s.fotosUseCase.Execute(

		&get_fotos_usecase.FotosInputDTO{
			OnNewItem: func(
				foto get_fotos_usecase.FotoDTO,
			) {
				out.Send(
					NewFotoDTO(
						foto.Id, foto.FotoUrl, foto.FotoUrl2, foto.Nome,
						foto.Partido, foto.UF, foto.Url,
						&pb.Redes{
							Red: foto.RedeSocial,
						},
					),
				)
			},
		},
	)

	if FotosOutputDTO.Error != nil {
		return FotosOutputDTO.Error
	}

	return nil
}
