package populate_db_usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/dariocasas/deputados3/src/configs"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
	"github.com/dariocasas/deputados3/src/internal/entity/deputado_index"
	"github.com/dariocasas/deputados3/src/internal/usecase/get_deputado_from_api_usecase"
	"github.com/dariocasas/deputados3/src/internal/usecase/scan_page"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
	"github.com/dariocasas/deputados3/src/pkg/http_client"

	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado_index"
)

type PopulateDbUseCase struct {
	config                  *configs.Config
	deputadoIndexRepository repo_deputado_index.DeputadoIndexRepositoryInterface
	deputadoRepository      repo_deputado.DeputadoRepositoryInterface
	deputadoApi             api_deputado.DeputadoApiInterface
	scanPageDeputado        scan_page.ScanPageDeputadoInterface
	httpClient              http_client.HttpClientInterface
	getDeputaFromApi        get_deputado_from_api_usecase.GetDeputadoFromApisUseCaseInterface
}

func NewPopulateDbUseCase(
	config *configs.Config,
	deputadoIndexRepository repo_deputado_index.DeputadoIndexRepositoryInterface,
	deputadoRepository repo_deputado.DeputadoRepositoryInterface,
	deputadoApi api_deputado.DeputadoApiInterface,
	scanPageDeputado scan_page.ScanPageDeputadoInterface,
	httpClient http_client.HttpClientInterface,
	getDeputaFromApi get_deputado_from_api_usecase.GetDeputadoFromApisUseCaseInterface,

) *PopulateDbUseCase {

	log.Println("NewPopulateDbUseCase")

	return &PopulateDbUseCase{
		config:                  config,
		deputadoIndexRepository: deputadoIndexRepository,
		deputadoRepository:      deputadoRepository,
		deputadoApi:             deputadoApi,
		scanPageDeputado:        scanPageDeputado,
		httpClient:              httpClient,
		getDeputaFromApi:        getDeputaFromApi,
	}
}

type resultDeputado = struct {
	Deputado    *deputado.Deputado
	Duration    time.Duration
	Error       error
	PartialTime PartialTime
	isApiValue  bool
}

func (p PopulateDbUseCase) Execute(input PopulateDbInputDTO) *PopulateDbOutputDTO {

	log.Printf("PopulateDbUseCase.Execute")

	itemsProcessed := 0

	elapsedTime := elapsed_time.NewElapsedTime()
	partialTime := PartialTime{}

	deputadoIndexRepositoryOutputDTO := p.deputadoIndexRepository.GetAll()
	if deputadoIndexRepositoryOutputDTO.Error != nil {
		log.Println(deputadoIndexRepositoryOutputDTO.Error)
		return &PopulateDbOutputDTO{
			Error:    deputadoIndexRepositoryOutputDTO.Error,
			Duration: elapsedTime.Elapsed(),
		}
	}

	if !p.deputadoRepository.CollectionExists() {
		err := p.deputadoRepository.CreateCollection()
		if err != nil {
			return &PopulateDbOutputDTO{
				Duration: elapsedTime.Elapsed(),
				Error:    err,
			}
		}
	}

	for _, idx := range deputadoIndexRepositoryOutputDTO.DeputadosIds {

		func(idx *deputado_index.DeputadoId) *PopulateDbOutputDTO {

			log.Printf("PopulateDbUseCase.Execute id:%d", idx.Id)
			elapsedTime.Reset()

			var errResult error
			var nameResult string
			var resultChan = make(chan *resultDeputado, 2)

			defer func() {
				input.OnNewItem(
					idx.Id,
					nameResult,
					errResult,
					elapsedTime,
					partialTime,
				)
			}()

			// ler deputado do banco de dados

			deputadoRepositoryOutputDTO := p.deputadoRepository.Find(idx.Id)

			deputado1 := deputadoRepositoryOutputDTO.Deputado
			deputadoExists := deputadoRepositoryOutputDTO.Error == nil
			if !deputadoExists {
				deputado1 = deputado.NewDeputado(idx.Id)
			}
			partialTime.DbRead = deputadoRepositoryOutputDTO.Duration

			// ler deputado pela API

			go func(resultChan chan *resultDeputado, partialTime PartialTime, id int) <-chan *resultDeputado {
				deputafromapi := p.getDeputaFromApi.Execute(
					get_deputado_from_api_usecase.NewGetDeputadoFromApiIntputDTO(
						id, input.TimeOut,
					),
				)
				partialTime.GetApi = deputafromapi.Duration
				resultChan <- &resultDeputado{
					Deputado:    deputafromapi.Deputado,
					Duration:    deputafromapi.Duration,
					Error:       deputafromapi.Error,
					PartialTime: partialTime,
					isApiValue:  true,
				}
				return resultChan
			}(resultChan, partialTime, deputado1.Id)

			go func(resultChan chan *resultDeputado, partialTime PartialTime, id int) <-chan *resultDeputado {

				// baixar html pagina deputado

				url := fmt.Sprintf("%s/%d", p.config.DeputadosPagesUrlBase, id)
				httpClientOutputDTO := p.httpClient.RequestPageWithRetry(
					&http_client.HttpClientWithRetryInputDTO{
						Url:     url,
						Timeout: input.TimeOut,
						MinLen:  10000,
						Retry:   3,
					},
				)
				partialTime.GetPage = httpClientOutputDTO.Duration
				if httpClientOutputDTO.Error != nil {

					resultChan <- &resultDeputado{
						Deputado:    deputado.NewDeputado(id),
						Duration:    httpClientOutputDTO.Duration,
						Error:       httpClientOutputDTO.Error,
						PartialTime: partialTime,
					}
					return resultChan
				}

				// escanear dados da pagina e atualizar deputado

				scanPageDeputadoOutputDTO := p.scanPageDeputado.Parse(
					scan_page.NewScanPageDeputadoInputDTO(
						id,
						string(*httpClientOutputDTO.Result),
					),
				)
				partialTime.ScanPage = scanPageDeputadoOutputDTO.Duration
				if scanPageDeputadoOutputDTO.Error != nil {

					resultChan <- &resultDeputado{
						Deputado:    deputado.NewDeputado(id),
						Duration:    scanPageDeputadoOutputDTO.Duration,
						Error:       scanPageDeputadoOutputDTO.Error,
						PartialTime: partialTime,
					}
					return resultChan
				}

				scanPageDeputadoOutputDTO.Deputado.Url = url

				resultChan <- &resultDeputado{
					Deputado:    scanPageDeputadoOutputDTO.Deputado,
					Duration:    scanPageDeputadoOutputDTO.Duration,
					Error:       scanPageDeputadoOutputDTO.Error,
					PartialTime: partialTime,
				}
				return resultChan
			}(resultChan, partialTime, deputado1.Id)

			res := copyValues(deputado1, <-resultChan, <-resultChan)
			partialTime = res.PartialTime
			deputado1 = res.Deputado

			// inserir no db

			databaseDeputadoInputDTO := repo_deputado.NewDatabaseDeputadoInputDTO(deputado1)

			if deputadoExists {
				databaseDeputadoOutputDTO := p.deputadoRepository.Update(databaseDeputadoInputDTO)
				if databaseDeputadoOutputDTO.Error != nil {
					errResult = databaseDeputadoOutputDTO.Error
					return &PopulateDbOutputDTO{
						Duration: elapsedTime.Elapsed(),
						Error:    errResult,
					}
				}
				partialTime.DbWrite = databaseDeputadoOutputDTO.Duration
			} else {
				databaseDeputadoOutputDTO := p.deputadoRepository.Insert(databaseDeputadoInputDTO)
				if databaseDeputadoOutputDTO.Error != nil {
					errResult = databaseDeputadoOutputDTO.Error
					return &PopulateDbOutputDTO{
						Duration: elapsedTime.Elapsed(),
						Error:    errResult,
					}
				}
				partialTime.DbWrite = databaseDeputadoOutputDTO.Duration
			}

			return &PopulateDbOutputDTO{
				Duration: elapsedTime.Elapsed(),
				Error:    nil,
			}

		}(idx)

		itemsProcessed++
	}

	return &PopulateDbOutputDTO{
		Duration:       elapsedTime.Elapsed(),
		Error:          nil,
		ItemsProcessed: itemsProcessed,
	}
}

func copyValues(deputado1 *deputado.Deputado, res1 *resultDeputado, res2 *resultDeputado) resultDeputado {

	result := resultDeputado{}
	var apiValues *resultDeputado
	var pageValues *resultDeputado

	if res1.isApiValue {
		apiValues = res1
		pageValues = res2
	} else {
		apiValues = res2
		pageValues = res1
	}

	result.Deputado = copyPageValues(deputado1, pageValues.Deputado)
	result.Deputado = copyApiValues(result.Deputado, apiValues.Deputado)

	result.PartialTime.GetApi = apiValues.PartialTime.GetApi
	result.PartialTime.GetPage = pageValues.PartialTime.GetPage
	result.PartialTime.ScanPage = pageValues.PartialTime.ScanPage

	return result
}

func copyApiValues(result *deputado.Deputado, deputado2 *deputado.Deputado) *deputado.Deputado {

	result.Nome = deputado2.UltimoStatus.Nome
	result.UF = deputado2.UltimoStatus.SiglaUf

	// result.Exercicio
	result.Partido = deputado2.UltimoStatus.SiglaPartido
	// result.Legislaturas
	result.NomeCivil = deputado2.NomeCivil
	// result.Email
	result.Uri = deputado2.Uri
	// result.Url
	// result.Foto = deputado2.UltimoStatus.UrlFoto
	// result.Partidos
	// result.Telefone
	// result.Endereco
	result.DataNasc = deputado2.DataNascimento
	result.Naturalidade = deputado2.MunicipioNascimento
	// result.Atuacoes
	// result.Presencas
	// result.Cargos
	// result.Frentes

	result.UltimoStatus = deputado2.UltimoStatus

	result.Cpf = deputado2.Cpf
	result.Sexo = deputado2.Sexo
	result.UrlWebsite = deputado2.UrlWebsite
	result.RedeSocial = deputado2.RedeSocial
	result.DataNascimento = deputado2.DataNascimento
	result.DataFalecimento = deputado2.DataFalecimento
	result.UfNascimento = deputado2.UfNascimento
	result.MunicipioNascimento = deputado2.MunicipioNascimento
	result.Escolaridade = deputado2.Escolaridade

	return result
}

func copyPageValues(result *deputado.Deputado, deputado2 *deputado.Deputado) *deputado.Deputado {

	result.Id = deputado2.Id
	result.Nome = deputado2.Nome
	result.NomeCivil = deputado2.NomeCivil
	result.Partido = deputado2.Partido
	result.UF = deputado2.UF

	result.UltimoStatus.Gabinete = deputado2.UltimoStatus.Gabinete

	result.UltimoStatus.UrlFoto2 = deputado2.UltimoStatus.UrlFoto2
	result.UltimoStatus.Email = deputado2.UltimoStatus.Email
	result.UltimoStatus.Gabinete.Telefone = deputado2.UltimoStatus.Gabinete.Telefone
	result.Endereco = deputado2.Endereco
	result.DataNasc = deputado2.DataNasc
	result.Naturalidade = deputado2.Naturalidade
	result.Atuacoes = deputado2.Atuacoes
	result.Presencas = deputado2.Presencas
	result.Cargos = deputado2.Cargos
	result.Frentes = deputado2.Frentes

	return result
}
