package deputado_index

import (
	"log"
)

type DeputadoId struct {
	Id int `json:"id" bson:"_id"`
}

type DeputadoIndex []*DeputadoId

func NewDeputadoIndex(ids DeputadoIndex) DeputadoIndex {
	log.Println("NewDeputadoIndex(ids...)")
	return ids
}
