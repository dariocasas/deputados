package grpc_service

import (
	"context"
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/infra/grpc/pb"
	"github.com/dariocasas/deputados3/src/internal/usecase/shutdown_usecase"
)

type ServerService struct {
	config *configs.Config
	pb.UnimplementedServerServiceServer
	shutdownUseCase shutdown_usecase.ShutdownUseCaseInterface
}

func NewServerService(
	config *configs.Config,
	shutdownUseCase shutdown_usecase.ShutdownUseCaseInterface,
) *ServerService {
	return &ServerService{
		config:          config,
		shutdownUseCase: shutdownUseCase,
	}
}

func (s *ServerService) Shutdown(ctx context.Context, in *pb.ShutdownRequest) (*pb.ShutdownResponse, error) {

	log.Println("ServerService.Shutdown()")

	if s.config.ServerAcceptShutdown {
		shutdownOutputDTO := s.shutdownUseCase.Execute(&shutdown_usecase.ShutdownInputDTO{})
		return &pb.ShutdownResponse{}, shutdownOutputDTO.Error
	}
	return &pb.ShutdownResponse{}, nil
}
