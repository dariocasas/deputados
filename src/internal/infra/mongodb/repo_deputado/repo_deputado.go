package repo_deputado

import (
	"context"
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeputadoRepository struct {
	client *mongo.Client
	config *configs.Config
}

func NewDeputadoRepository(
	client *mongo.Client,
	config *configs.Config,
) *DeputadoRepository {

	log.Println("NewDeputadoRepository")

	return &DeputadoRepository{
		client: client,
		config: config,
	}
}

func (r *DeputadoRepository) Find(id int) *DatabaseDeputadoOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	log.Printf("DeputadoRepository.Find(%d)", id)

	newDeputado := deputado.NewDeputado(id)

	coll := r.client.
		Database(r.config.DatabaseName).
		Collection(r.config.DeputadosCollection)

	filter := bson.D{{Key: "id", Value: id}}

	result := coll.FindOne(context.TODO(), filter)
	err := result.Decode(&newDeputado)
	if err != nil {
		log.Println(err)
		return NewDatabaseDeputadoOutputDTO(nil, elapsedTime.Elapsed(), err)
	}

	return NewDatabaseDeputadoOutputDTO(newDeputado, elapsedTime.Elapsed(), nil)
}

func (r *DeputadoRepository) FindAll() *DatabaseDeputadosOutputDTO {
	log.Println("DeputadoRepository.FindAll())")

	elapsedTime := elapsed_time.NewElapsedTime()

	deputados := []*deputado.Deputado{}

	coll := r.client.
		Database(r.config.DatabaseName).
		Collection(r.config.DeputadosCollection)

	sort := bson.D{{Key: "nome", Value: 1}}
	opts := options.Find().SetSort(sort)

	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return NewDatabaseDeputadosOutputDTO(nil, elapsedTime.Elapsed(), err)
	}

	if err = cursor.All(context.TODO(), &deputados); err != nil {
		log.Println(err)
	}

	return NewDatabaseDeputadosOutputDTO(deputados, elapsedTime.Elapsed(), nil)
}

func (r *DeputadoRepository) Insert(inputDTO DatabaseDeputadoInputDTO) *DatabaseDeputadoOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	log.Println("DeputadoRepository.Insert(deputado)")

	coll := r.client.
		Database(r.config.DatabaseName).
		Collection(r.config.DeputadosCollection)

	_, err := coll.InsertOne(context.TODO(), inputDTO)
	if err != nil {
		log.Println(err)
		return NewDatabaseDeputadoOutputDTO(nil, elapsedTime.Elapsed(), err)
	}
	return NewDatabaseDeputadoOutputDTO(nil, elapsedTime.Elapsed(), nil)
}

func (r *DeputadoRepository) Update(inputDTO DatabaseDeputadoInputDTO) *DatabaseDeputadoOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	log.Printf("DeputadoRepository.Update(%d)", inputDTO.Id)

	coll := r.client.
		Database(r.config.DatabaseName).
		Collection(r.config.DeputadosCollection)

	filter := bson.D{{Key: "id", Value: inputDTO.Id}}
	_, err := coll.ReplaceOne(context.TODO(), filter, inputDTO)
	if err != nil {
		log.Println(err)
		return NewDatabaseDeputadoOutputDTO(nil, elapsedTime.Elapsed(), err)
	}
	return NewDatabaseDeputadoOutputDTO(nil, elapsedTime.Elapsed(), nil)
}

// Collection related

func (r *DeputadoRepository) CollectionExists() bool {
	return mongodb.CollectionExists(
		r.client.Database(r.config.DatabaseName),
		r.config.DeputadosCollection)
}

func (r *DeputadoRepository) DropCollention() error {
	log.Println("DeputadoRepository.DropCollention())")
	return r.client.Database(r.config.DatabaseName).
		Collection(r.config.DeputadosCollection).
		Drop(context.TODO())
}

func (r *DeputadoRepository) CreateCollection() error {
	log.Printf("DeputadoRepository.CreateCollection")

	coll := r.client.
		Database(r.config.DatabaseName).
		Collection(r.config.DeputadosCollection)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "nome", Value: 1}},
		Options: options.Index().SetName("idx_nome").SetUnique(true),
	}
	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}

	indexModel = mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetName("idx_id"),
	}
	_, err = coll.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

// Cursor related

func (r *DeputadoRepository) GetCursor() (any, error) {
	log.Println("DeputadoRepository.GetCursor())")

	coll := r.client.
		Database(r.config.DatabaseName).
		Collection(r.config.DeputadosCollection)

	sort := bson.D{{Key: "nome", Value: 1}}
	opts := options.Find().SetSort(sort)

	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func (r *DeputadoRepository) ReadCursor(cursor any, out chan *deputado.Deputado) error {

	defer close(out)

	cur := cursor.(*mongo.Cursor)
	for cur.Next(context.TODO()) {
		var result *deputado.Deputado
		if err := cur.Decode(&result); err != nil {
			return err
		}
		out <- result
	}
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}
func (r *DeputadoRepository) CloseCursor(cursor any) {
	defer cursor.(*mongo.Cursor).Close(context.TODO())
}
