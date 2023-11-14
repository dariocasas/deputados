package scan_page_test

import (
	"testing"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
	"github.com/dariocasas/deputados3/src/internal/usecase/scan_page"

	"github.com/dariocasas/deputados3/src/pkg/class_finder"
	"github.com/dariocasas/deputados3/src/pkg/http_client"
	"github.com/stretchr/testify/assert"
)

func TestScanPageDeputado(t *testing.T) {

	config, err := configs.NewConfig("c:\\Users\\dario\\dev\\deputados3\\src\\cmd\\server")
	if err != nil {
		panic(err)
	}
	deputado220623 := deputado.NewDeputado(220623)
	httpClient := http_client.NewHttpClient(config)
	classFinder := class_finder.NewClassFinder()
	scanPageDeputado := scan_page.NewScanPageDeputado(classFinder)
	httpClientOutputDTO := httpClient.RequestPage(
		&http_client.HttpClientInputDTO{
			Url:     "https://www.camara.leg.br/deputados/220623",
			Timeout: 15000,
		},
	)
	assert.Nil(t, httpClientOutputDTO.Error)

	scanPageDeputadoOutputDTO := scanPageDeputado.Parse(
		scan_page.NewScanPageDeputadoInputDTO(
			deputado220623.Id,
			string(*httpClientOutputDTO.Result),
		),
	)
	assert.Nil(t, scanPageDeputadoOutputDTO.Error)

	want := "https://www.camara.leg.br/internet/deputado/bandep/pagina_do_deputado/220623.jpg"

	has := scanPageDeputadoOutputDTO.Deputado.UltimoStatus.UrlFoto2
	assert.Equal(t, want, has, "result does not match")

}
