package shutdown_usecase

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dariocasas/deputados3/src/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type ShutdownUseCase struct {
	grpc   *grpc.Server
	client *mongo.Client
	config *configs.Config
}

func NewShutdownUseCase(
	grpc *grpc.Server,
	client *mongo.Client,
	config *configs.Config,
) *ShutdownUseCase {

	log.Println("NewShutdownUseCase")

	return &ShutdownUseCase{
		grpc:   grpc,
		client: client,
		config: config,
	}
}

func (p ShutdownUseCase) Execute(input *ShutdownInputDTO) *ShutdownOutputDTO {

	log.Println("ShutdownUseCase.Execute()")

	err := p.client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
	}

	log.Println("responsing to shutdown request")

	go func() {
		log.Println("exiting in 1 sec")
		time.Sleep(time.Second)
		p.grpc.Stop()
		log.Println("exiting")
		os.Exit(0)
	}()

	return &ShutdownOutputDTO{Error: err}

}
