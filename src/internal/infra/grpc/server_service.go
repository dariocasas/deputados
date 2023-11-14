package grpc_service

import (
	"context"

	"github.com/dariocasas/deputados3/src/internal/infra/grpc/pb"
	"github.com/dariocasas/deputados3/src/internal/usecase/shutdown_usecase"
)

type ServerService struct {
	pb.UnimplementedServerServiceServer
	shutdownUseCase shutdown_usecase.ShutdownUseCaseInterface
}

func NewServerService(
	shutdownUseCase shutdown_usecase.ShutdownUseCaseInterface,
) *ServerService {
	return &ServerService{shutdownUseCase: shutdownUseCase}
}

func (s *ServerService) Shutdown(ctx context.Context, in *pb.ShutdownRequest) (*pb.ShutdownResponse, error) {
	shutdownOutputDTO := s.shutdownUseCase.Execute(&shutdown_usecase.ShutdownInputDTO{})
	return &pb.ShutdownResponse{}, shutdownOutputDTO.Error
}
