package http_client_test

import (
	"context"
	"testing"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/pkg/http_client"
	"github.com/stretchr/testify/assert"
)

func TestRequestJsonTimeout(t *testing.T) {
	config := &configs.Config{
		HttpClientTimeOutMilliSecs: 10000,
	}

	want := `{"dados":{"id":220623,"uri":"https://dadosabertos.camara.leg.br/api/v2/deputados/220623","nomeCivil":"DUDA SALABERT ROSA","ultimoStatus":{"id":220623,"uri":"https://dadosabertos.camara.leg.br/api/v2/deputados/220623","nome":"Duda Salabert","siglaPartido":"PDT","uriPartido":"https://dadosabertos.camara.leg.br/api/v2/partidos/36786","siglaUf":"MG","idLegislatura":57,"urlFoto":"https://www.camara.leg.br/internet/deputado/bandep/220623.jpg","email":"dep.dudasalabert@camara.leg.br","data":"2023-02-01T12:05","nomeEleitoral":"Duda Salabert","gabinete":{"nome":"840","predio":"4","sala":"840","andar":"8","telefone":"3215-5840","email":"dep.dudasalabert@camara.leg.br"},"situacao":"Exercício","condicaoEleitoral":"Titular","descricaoStatus":null},"cpf":"04967383645","sexo":"F","urlWebsite":null,"redeSocial":["https://twitter.com/DudaSalabert","https://www.facebook.com/DudaSalabert","https://www.instagram.com/duda_salabert"],"dataNascimento":"1981-05-02","dataFalecimento":null,"ufNascimento":"MG","municipioNascimento":"Belo Horizonte","escolaridade":"Superior"},"links":[{"rel":"self","href":"https://dadosabertos.camara.leg.br/api/v2/deputados/220623"}]}`
	url := "https://dadosabertos.camara.leg.br/api/v2/deputados/220623"
	has := http_client.NewHttpClient(config).RequestJson(
		&http_client.HttpClientInputDTO{
			Url: url,
		},
	)

	assert.Nil(t, has.Error)
	assert.Equal(t, want, string(*has.Result), "result does not match")

	config = &configs.Config{
		HttpClientTimeOutMilliSecs: 0,
	}
	has = http_client.NewHttpClient(config).RequestJson(
		&http_client.HttpClientInputDTO{
			Url:     url,
			Timeout: 10000,
		},
	)

	assert.Nil(t, has.Error)
	assert.Equal(t, want, string(*has.Result), "result does not match")

	config = &configs.Config{
		HttpClientTimeOutMilliSecs: 0,
	}
	has = http_client.NewHttpClient(config).RequestJson(
		&http_client.HttpClientInputDTO{
			Url: url,
		},
	)

	assert.Nil(t, has.Result)
	assert.Contains(t, has.Error.Error(), context.DeadlineExceeded.Error(), "timeout")
}

func TestRequestJson(t *testing.T) {
	config := &configs.Config{
		HttpClientTimeOutMilliSecs: 10000,
	}

	want := "400 "
	url := "https://dadosabertos.camara.leg.br/api/v2/deputados/220623erro"
	has := http_client.NewHttpClient(config).RequestJson(
		&http_client.HttpClientInputDTO{
			Url: url,
		},
	)

	assert.Nil(t, has.Result)
	assert.Contains(t, has.Error.Error(), want)
}
