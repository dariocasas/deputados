package db_exists_usecase

import (
	"context"
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado_index"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbExistsUseCase struct {
	client  *mongo.Client
	config  *configs.Config
	repoids repo_deputado_index.DeputadoIndexRepositoryInterface
}

func NewDbStatusUseCase(
	client *mongo.Client,
	config *configs.Config,
	repoids repo_deputado_index.DeputadoIndexRepositoryInterface,
) *DbExistsUseCase {

	log.Println("NewDbExistsUseCase")

	return &DbExistsUseCase{
		client:  client,
		config:  config,
		repoids: repoids,
	}
}

func (u DbExistsUseCase) Execute() *DbStatusUseCaseOutputDTO {

	log.Println("DbExistsUseCase.Execute()")

	db := u.client.Database(u.config.DatabaseName)

	collNames, err := db.ListCollectionNames(
		context.TODO(),
		bson.D{},
	)
	if err != nil {
		return NewDbStatusUseCaseOutputDTO(0, 0, err)
	}

	hasIndex := Contains(collNames, u.config.DeputadosIdsCollection)
	if !hasIndex {
		return NewDbStatusUseCaseOutputDTO(0, 0, nil)
	}

	output := u.repoids.GetAll()

	hasDeputados := Contains(collNames, u.config.DeputadosCollection)

	if !hasDeputados {
		return NewDbStatusUseCaseOutputDTO(1, int32(len(output.DeputadosIds)), nil)
	}
	return NewDbStatusUseCaseOutputDTO(2, int32(len(output.DeputadosIds)), nil)
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
