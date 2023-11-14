package api_deputados_all

import (
	"log"
	"time"

	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputados_all/model"
)

type ApiDeputadosAllInterface interface {
	GetAll(input *ApiDeputadosAllInputDTO) *ApiDeputadosAllOutputDTO
}

type ApiDeputadosAllInputDTO struct {
	timeout  int
	maxItems int
}

func NewApiDeputadosAllInputDTO(timeout int, maxItems int) *ApiDeputadosAllInputDTO {
	log.Println("NewApiDeputadosAllInputDTO")
	return &ApiDeputadosAllInputDTO{
		timeout:  timeout,
		maxItems: maxItems,
	}
}

type ApiDeputadosAllOutputDTO struct {
	Deputados *model.Deputados
	Duration  time.Duration
	Error     error
}

func NewApiDeputadoIdOutputDTO(ids *model.Deputados, duration time.Duration, error error) *ApiDeputadosAllOutputDTO {
	log.Println("NewApiDeputadoIdOutputDTO")
	return &ApiDeputadosAllOutputDTO{
		Deputados: ids,
		Duration:  duration,
		Error:     error,
	}
}
