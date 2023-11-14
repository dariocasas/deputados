package repo_deputado_index

import (
	"log"
	"time"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado_index"
)

type DeputadoIndexRepositoryInterface interface {
	InsertAll(ids DeputadoIndexRepositoryInputDTO) *DeputadoIndexRepositoryOutputDTO
	GetAll() *DeputadoIndexRepositoryOutputDTO
	CollectionExists() bool
	DropCollention() error
}

// --------------------- INPUT DTO ---------------------

type DeputadoIndexRepositoryInputDTO = deputado_index.DeputadoIndex

func NewDeputadoIndexRepositoryInputDTO(ids deputado_index.DeputadoIndex) DeputadoIndexRepositoryInputDTO {
	log.Println("NewDatabaseDeputadoIdInputDTO(ids...)")
	return DeputadoIndexRepositoryInputDTO(ids)
}

// --------------------- OUTPUT DTO --------------------

type DeputadoIndexRepositoryOutputDTO struct {
	DeputadosIds deputado_index.DeputadoIndex
	Duration     time.Duration
	Error        error
}

func NewDeputadoIndexRepositoryOutputDTO(ids deputado_index.DeputadoIndex, duration time.Duration, error error) *DeputadoIndexRepositoryOutputDTO {
	log.Println("NewDatabaseDeputadoIdOutputDTO")
	return &DeputadoIndexRepositoryOutputDTO{
		DeputadosIds: ids,
		Duration:     duration,
		Error:        error,
	}
}
