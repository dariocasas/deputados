package populate_index_usecase

import (
	"time"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado_index"
)

type PopulateIndexInputDTO struct {
	maxItems int
	list     deputado_index.DeputadoIndex
}

func NewPopulateIndexInputDTO(
	maxItems int,
	list deputado_index.DeputadoIndex,
) *PopulateIndexInputDTO {
	return &PopulateIndexInputDTO{
		maxItems: maxItems,
		list:     list,
	}
}

type PopulateIndexOutputDTO = struct {
	Duration time.Duration
	Error    error
}

type PopulateIndexUseCaseInterface interface {
	Execute(input *PopulateIndexInputDTO) *PopulateIndexOutputDTO
}
