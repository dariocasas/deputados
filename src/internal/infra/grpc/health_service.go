package grpc_service

import (
	"context"
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/infra/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HealthService struct {
	pb.UnimplementedHealthServiceServer
	config *configs.Config
}

func NewHealthService(
	config *configs.Config,
) *HealthService {

	healthService := &HealthService{
		config: config,
	}
	return healthService
}

func (s *HealthService) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {

	log.Println("Health.Check()")

	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthService) Watch(in *pb.HealthCheckRequest, out pb.HealthService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
