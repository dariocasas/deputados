package fetch_deputados_ids_usecase

import (
	"time"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado_index"
)

type FetchDeputadosIdsUseCaseInterface interface {
	Execute(input *FetchDeputadosIdsInputDTO) *FetchDeputadosIdsOutputDTO
}

type FetchDeputadosIdsInputDTO struct {
	timeout  int
	maxItems int
}

func NewFetchDeputadosIdsInputDTO(
	timeout int,
	maxItems int,
) *FetchDeputadosIdsInputDTO {
	return &FetchDeputadosIdsInputDTO{
		timeout:  timeout,
		maxItems: maxItems,
	}
}

type FetchDeputadosIdsOutputDTO = struct {
	List     deputado_index.DeputadoIndex
	Duration time.Duration
	Error    error
}

func NewFetchDeputadosIdsOutputDTO(
	List deputado_index.DeputadoIndex,
	Duration time.Duration,
	Error error,
) *FetchDeputadosIdsOutputDTO {
	return &FetchDeputadosIdsOutputDTO{
		List:     List,
		Duration: Duration,
		Error:    Error,
	}
}
