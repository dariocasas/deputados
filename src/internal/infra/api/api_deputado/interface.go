package api_deputado

import (
	"time"

	model "github.com/dariocasas/deputados3/src/internal/infra/api/api_deputado/model"
)

type DeputadoApiInterface interface {
	Get(input *ApiDeputadoIntputDTO) *ApiDeputadoOutputDTO
}

type ApiDeputadoIntputDTO struct {
	id      int
	timeout int
}

func NewApiDeputadoIntputDTO(id int, timeout int) *ApiDeputadoIntputDTO {
	return &ApiDeputadoIntputDTO{
		id:      id,
		timeout: timeout,
	}
}

type ApiDeputadoOutputDTO struct {
	Deputado *model.Deputado
	Duration time.Duration
	Error    error
}

func NewApiDeputadoOutputDTO(deputado *model.Deputado, duration time.Duration, error error) *ApiDeputadoOutputDTO {
	return &ApiDeputadoOutputDTO{
		Deputado: deputado,
		Duration: duration,
		Error:    error,
	}
}
