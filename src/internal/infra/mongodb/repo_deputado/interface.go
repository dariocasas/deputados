package repo_deputado

import (
	"time"

	entity "github.com/dariocasas/deputados3/src/internal/entity/deputado"
)

type DeputadoRepositoryInterface interface {
	Find(id int) *DatabaseDeputadoOutputDTO
	Insert(deputado DatabaseDeputadoInputDTO) *DatabaseDeputadoOutputDTO
	Update(deputado DatabaseDeputadoInputDTO) *DatabaseDeputadoOutputDTO
	FindAll() *DatabaseDeputadosOutputDTO

	CollectionExists() bool
	DropCollention() error
	CreateCollection() error

	GetCursor() (any, error)
	ReadCursor(cursor any, out chan *entity.Deputado) error
	CloseCursor(cursor any)
}

// --------------------- INPUT DTO ---------------------

type DatabaseDeputadoInputDTO = *entity.Deputado

func NewDatabaseDeputadoInputDTO(deputado *entity.Deputado) DatabaseDeputadoInputDTO {
	return DatabaseDeputadoInputDTO(deputado)
}

// --------------------- OUTPUT DTO --------------------

type DatabaseDeputadoOutputDTO struct {
	Deputado *entity.Deputado
	Duration time.Duration
	Error    error
}

func NewDatabaseDeputadoOutputDTO(deputado *entity.Deputado, duration time.Duration, error error) *DatabaseDeputadoOutputDTO {
	return &DatabaseDeputadoOutputDTO{
		Deputado: deputado,
		Duration: duration,
		Error:    error,
	}
}

type DatabaseDeputadosOutputDTO struct {
	Deputados []*entity.Deputado
	Duration  time.Duration
	Error     error
}

func NewDatabaseDeputadosOutputDTO(deputados []*entity.Deputado, duration time.Duration, error error) *DatabaseDeputadosOutputDTO {
	return &DatabaseDeputadosOutputDTO{
		Deputados: deputados,
		Duration:  duration,
		Error:     error,
	}
}
