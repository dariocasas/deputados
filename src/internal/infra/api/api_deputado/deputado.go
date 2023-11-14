package api_deputado

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dariocasas/deputados3/src/configs"

	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputado/model"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
	"github.com/dariocasas/deputados3/src/pkg/http_client"
)

type ApiDeputado struct {
	config     *configs.Config
	httpclient *http_client.HttpClient
}

func NewApiDeputado(
	config *configs.Config,
	httpclient *http_client.HttpClient,
) *ApiDeputado {
	log.Println("NewApiDeputado")

	return &ApiDeputado{
		config:     config,
		httpclient: httpclient,
	}
}

type Wrapper struct {
	Dados *model.Deputado `json:"dados"`
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

func (r *ApiDeputado) Get(input *ApiDeputadoIntputDTO) *ApiDeputadoOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	var wrapper Wrapper

	url := fmt.Sprintf("%sdeputados/%d", r.config.ApiBaseUrl, input.id)
	log.Printf("ApiDeputado.Find(%d) %s", input.id, url)

	jsonBytes := r.httpclient.RequestJson(
		&http_client.HttpClientInputDTO{
			Url:     url,
			Timeout: input.timeout,
		},
	)
	if jsonBytes.Error != nil {
		log.Printf("Error requesting %s: %s\n", url, jsonBytes.Error)
		return &ApiDeputadoOutputDTO{
			Duration: elapsedTime.Elapsed(),
			Error:    jsonBytes.Error,
		}
	}

	err := json.Unmarshal(*jsonBytes.Result, &wrapper)
	if err != nil {
		return &ApiDeputadoOutputDTO{
			Duration: elapsedTime.Elapsed(),
			Error:    err,
		}
	}

	return NewApiDeputadoOutputDTO(wrapper.Dados, elapsedTime.Elapsed(), nil)
}
