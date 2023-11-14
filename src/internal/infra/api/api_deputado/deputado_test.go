package api_deputado_test

import (
	"testing"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputado/model"

	"github.com/dariocasas/deputados3/src/pkg/http_client"
	"github.com/stretchr/testify/assert"
)

func TestDeputado(t *testing.T) {

	want := &model.Deputado{
		Id:        220623,
		Uri:       "https://dadosabertos.camara.leg.br/api/v2/deputados/220623",
		NomeCivil: "DUDA SALABERT ROSA",
		UltimoStatus: &deputado.UltimoStatus{
			Id:            220623,
			Uri:           "https://dadosabertos.camara.leg.br/api/v2/deputados/220623",
			Nome:          "Duda Salabert",
			SiglaPartido:  "PDT",
			UriPartido:    "https://dadosabertos.camara.leg.br/api/v2/partidos/36786",
			SiglaUf:       "MG",
			IdLegislatura: 57,
			UrlFoto:       "https://www.camara.leg.br/internet/deputado/bandep/220623.jpg",
			Email:         "dep.dudasalabert@camara.leg.br",
			Data:          "2023-02-01T12:05",
			NomeEleitoral: "Duda Salabert",
			Gabinete: &deputado.Gabinete{
				Nome:     "840",
				Predio:   "4",
				Sala:     "840",
				Andar:    "8",
				Telefone: "3215-5840",
				Email:    "dep.dudasalabert@camara.leg.br",
			},
			Situacao:          "Exercício",
			CondicaoEleitoral: "Titular",
			// DescricaoStatus:   null,
		},
		Cpf:        "04967383645",
		Sexo:       "F",
		UrlWebsite: "",
		RedeSocial: []string{
			"https://twitter.com/DudaSalabert",
			"https://www.facebook.com/DudaSalabert",
			"https://www.instagram.com/duda_salabert",
		},
		DataNascimento:      "1981-05-02",
		DataFalecimento:     "",
		UfNascimento:        "MG",
		MunicipioNascimento: "Belo Horizonte",
		Escolaridade:        "Superior",
	}

	config, _ := configs.NewConfig("c:\\Users\\dario\\dev\\deputados3\\src\\cmd\\server")
	httpClient := http_client.NewHttpClient(config)
	deputadoApi := api_deputado.NewApiDeputado(config, httpClient)

	apiDeputadoOutputDTO := deputadoApi.Get(api_deputado.NewApiDeputadoIntputDTO(220623, 15000))
	assert.Nil(t, apiDeputadoOutputDTO.Error)
	has := apiDeputadoOutputDTO.Deputado
	assert.Equal(t, want, has, "result does not match")

}
