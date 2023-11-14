package repo_deputado_index

import (
	"context"
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/entity/deputado_index"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeputadoIndexRepository struct {
	client *mongo.Client
	config *configs.Config
}

func NewDeputadoIndexRepository(
	client *mongo.Client,
	config *configs.Config,
) *DeputadoIndexRepository {

	log.Println("NewDeputadoIndexRepository")

	return &DeputadoIndexRepository{
		client: client,
		config: config,
	}
}

func (r *DeputadoIndexRepository) InsertAll(ids DeputadoIndexRepositoryInputDTO) *DeputadoIndexRepositoryOutputDTO {

	log.Println("DeputadoIndexRepository.InsertAll(ids...)")

	elapsedTime := elapsed_time.NewElapsedTime()

	var data []any
	for _, v := range ids {
		data = append(data, v)
	}

	db := r.client.
		Database(r.config.DatabaseName)

	createIndex := mongodb.CollectionExists(db, r.config.DeputadosIdsCollection)

	coll := r.client.
		Database(r.config.DatabaseName).
		Collection(r.config.DeputadosIdsCollection)

	if createIndex {
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "_id", Value: 1}},
			Options: options.Index().SetName("idx"),
		}

		_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Println(err)
			return NewDeputadoIndexRepositoryOutputDTO(
				nil, elapsedTime.Elapsed(), err,
			)
		}

	}

	_, err := coll.InsertMany(context.TODO(), data)
	if err != nil {
		log.Println(err)
		return NewDeputadoIndexRepositoryOutputDTO(
			nil, elapsedTime.Elapsed(), err,
		)
	}

	return NewDeputadoIndexRepositoryOutputDTO(
		nil, elapsedTime.Elapsed(), nil,
	)
}

func (r *DeputadoIndexRepository) GetAll() *DeputadoIndexRepositoryOutputDTO {

	log.Println("DeputadoIndexRepository.GetAll()")

	elapsedTime := elapsed_time.NewElapsedTime()
	ids := []*deputado_index.DeputadoId{}

	coll := r.client.
		Database(r.config.DatabaseName).
		Collection(r.config.DeputadosIdsCollection)

	sort := bson.D{{Key: "_id", Value: 1}}
	opts := options.Find().SetSort(sort)

	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return NewDeputadoIndexRepositoryOutputDTO(nil, elapsedTime.Elapsed(), err)
	}
	if err = cursor.All(context.TODO(), &ids); err != nil {
		log.Println(err)
	}

	return NewDeputadoIndexRepositoryOutputDTO(deputado_index.NewDeputadoIndex(ids), elapsedTime.Elapsed(), nil)
}

func (r *DeputadoIndexRepository) CollectionExists() bool {
	return mongodb.CollectionExists(
		r.client.Database(r.config.DatabaseName),
		r.config.DeputadosIdsCollection)
}

func (r *DeputadoIndexRepository) DropCollention() error {
	log.Println("DeputadoIndexRepository.DropCollention())")
	return r.client.Database(r.config.DatabaseName).
		Collection(r.config.DeputadosIdsCollection).
		Drop(context.TODO())
}
