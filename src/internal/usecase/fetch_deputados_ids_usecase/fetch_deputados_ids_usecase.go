package fetch_deputados_ids_usecase

import (
	"log"

	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputados_all"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado_index"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
)

type FetchDeputadosIdsUseCase struct {
	DeputadoIdRepository repo_deputado_index.DeputadoIndexRepositoryInterface
	ApiDeputadosAll      api_deputados_all.ApiDeputadosAllInterface
}

func NewFetchDeputadosIdsUseCase(
	deputadoIndexRepository repo_deputado_index.DeputadoIndexRepositoryInterface,
	ApiDeputadosAll api_deputados_all.ApiDeputadosAllInterface,

) *FetchDeputadosIdsUseCase {

	log.Println("NewFetchDeputadosIdsUseCase")

	return &FetchDeputadosIdsUseCase{
		DeputadoIdRepository: deputadoIndexRepository,
		ApiDeputadosAll:      ApiDeputadosAll,
	}
}

func (u FetchDeputadosIdsUseCase) Execute(input *FetchDeputadosIdsInputDTO) *FetchDeputadosIdsOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	log.Println("FetchDeputadosIdsUseCase.Execute()")

	apiDeputadoIdOutputDTO := u.ApiDeputadosAll.GetAll(
		api_deputados_all.NewApiDeputadosAllInputDTO(
			input.timeout,
			input.maxItems,
		),
	)
	if apiDeputadoIdOutputDTO.Error != nil {
		return NewFetchDeputadosIdsOutputDTO(
			nil,
			elapsedTime.Elapsed(),
			apiDeputadoIdOutputDTO.Error,
		)
	}

	return NewFetchDeputadosIdsOutputDTO(
		apiDeputadoIdOutputDTO.Deputados.Ids(),
		elapsedTime.Elapsed(),
		nil,
	)
}
