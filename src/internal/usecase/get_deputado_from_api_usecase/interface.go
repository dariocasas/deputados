package get_deputado_from_api_usecase

import (
	"time"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
)

type GetDeputadoFromApisUseCaseInterface interface {
	Execute(deputado *GetDeputadoFromApiInputDTO) *GetDeputadoFromApiOutputDTO
}

type GetDeputadoFromApiInputDTO struct {
	id      int
	timeout int
}

func NewGetDeputadoFromApiIntputDTO(
	id int,
	timeout int,
) *GetDeputadoFromApiInputDTO {
	return &GetDeputadoFromApiInputDTO{
		id:      id,
		timeout: timeout,
	}
}

type GetDeputadoFromApiOutputDTO = struct {
	Deputado *deputado.Deputado
	Duration time.Duration
	Error    error
}

func NewGetDeputadoFromApiOutputDTO(
	Deputado *deputado.Deputado,
	Duration time.Duration,
	Error error,
) *GetDeputadoFromApiOutputDTO {
	return &GetDeputadoFromApiOutputDTO{
		Deputado: Deputado,
		Duration: Duration,
		Error:    Error,
	}
}
