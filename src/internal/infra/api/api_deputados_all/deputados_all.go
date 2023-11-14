package api_deputados_all

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	model "github.com/dariocasas/deputados3/src/internal/infra/api/api_deputados_all/model"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
	"github.com/dariocasas/deputados3/src/pkg/http_client"
)

type ApiDeputadosAll struct {
	config     *configs.Config
	httpClient http_client.HttpClientInterface
}

func NewApiDeputadosAll(
	config *configs.Config,
	httpClient http_client.HttpClientInterface,
) *ApiDeputadosAll {
	log.Println("NewApiDeputadosAll")

	return &ApiDeputadosAll{
		config:     config,
		httpClient: httpClient,
	}
}

type Wrapper struct {
	Dados []*model.Deputado `json:"dados"`
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

func (r *ApiDeputadosAll) GetAll(input *ApiDeputadosAllInputDTO) *ApiDeputadosAllOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	var wrapper Wrapper

	queryStr := "deputados"
	if input.maxItems > 0 {
		queryStr = fmt.Sprintf("deputados?itens=%d", input.maxItems)
	}

	url := fmt.Sprintf("%s%s", r.config.ApiBaseUrl, queryStr)
	log.Printf("ApiDeputadosAll.GetAll() %s", url)

	httpClientOutputDTO := r.httpClient.RequestJson(
		&http_client.HttpClientInputDTO{
			Url:     url,
			Timeout: input.timeout,
		})
	if httpClientOutputDTO.Error != nil {
		log.Printf("Error requesting %s: %s\n", url, httpClientOutputDTO.Error)
		return NewApiDeputadoIdOutputDTO(
			nil, elapsedTime.Elapsed(), httpClientOutputDTO.Error,
		)
	}

	err := json.Unmarshal(*httpClientOutputDTO.Result, &wrapper)
	if err != nil {
		log.Println(err)
		return NewApiDeputadoIdOutputDTO(nil, elapsedTime.Elapsed(), err)
	}

	deps, err := model.NewDeputados(wrapper.Dados)
	if err != nil {
		log.Println(err)
		return NewApiDeputadoIdOutputDTO(nil, elapsedTime.Elapsed(), err)
	}

	return NewApiDeputadoIdOutputDTO(&deps, elapsedTime.Elapsed(), nil)

}
