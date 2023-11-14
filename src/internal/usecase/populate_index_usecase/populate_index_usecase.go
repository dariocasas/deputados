package populate_index_usecase

import (
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputados_all"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado_index"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
)

type PopulateIndexUseCase struct {
	config               *configs.Config
	DeputadoIdRepository repo_deputado_index.DeputadoIndexRepositoryInterface
	ApiDeputadosAll      api_deputados_all.ApiDeputadosAllInterface
}

func NewPopulateIndexUseCase(
	config *configs.Config,
	deputadoIndexRepository repo_deputado_index.DeputadoIndexRepositoryInterface,
	ApiDeputadosAll api_deputados_all.ApiDeputadosAllInterface,
) *PopulateIndexUseCase {

	log.Println("NewPopulateIndexUseCase")

	return &PopulateIndexUseCase{
		DeputadoIdRepository: deputadoIndexRepository,
		ApiDeputadosAll:      ApiDeputadosAll,
		config:               config,
	}
}

func (u PopulateIndexUseCase) Execute(input *PopulateIndexInputDTO) *PopulateIndexOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	log.Println("PopulateIndexUseCase.Execute()")

	databaseDeputadoIdInputDTO := repo_deputado_index.NewDeputadoIndexRepositoryInputDTO(
		input.list,
	)

	if u.DeputadoIdRepository.CollectionExists() {
		u.DeputadoIdRepository.DropCollention()
	}

	apiDeputadosAllOutputDTO := u.ApiDeputadosAll.GetAll(
		api_deputados_all.NewApiDeputadosAllInputDTO(
			u.config.HttpClientTimeOutMilliSecs,
			input.maxItems,
		),
	)
	if apiDeputadosAllOutputDTO.Error != nil {
		return &PopulateIndexOutputDTO{
			Error:    apiDeputadosAllOutputDTO.Error,
			Duration: elapsedTime.Elapsed(),
		}
	}

	deputadoIndexRepositoryOutputDTO := u.DeputadoIdRepository.InsertAll(databaseDeputadoIdInputDTO)
	if deputadoIndexRepositoryOutputDTO.Error != nil {
		return &PopulateIndexOutputDTO{
			Error:    deputadoIndexRepositoryOutputDTO.Error,
			Duration: elapsedTime.Elapsed(),
		}
	}
	return &PopulateIndexOutputDTO{
		Duration: elapsedTime.Elapsed(),
		Error:    nil,
	}
}
