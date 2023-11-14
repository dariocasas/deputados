package get_deputado_from_api_usecase

import (
	"log"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/api/api_deputado"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
)

type GetDeputaFromApidoUseCase struct {
	DeputadoApi api_deputado.DeputadoApiInterface
}

func NewGetDeputaFromApiUseCase(
	deputadoApi api_deputado.DeputadoApiInterface,
) *GetDeputaFromApidoUseCase {

	log.Println("NewGetDeputadoFromApiUseCase")

	return &GetDeputaFromApidoUseCase{
		DeputadoApi: deputadoApi,
	}
}

func (u *GetDeputaFromApidoUseCase) Execute(input *GetDeputadoFromApiInputDTO) *GetDeputadoFromApiOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	log.Println("GetDeputadoFromApiUseCase.Execute()")

	result := deputado.NewDeputado(input.id)

	apiDeputadoOutputDTO := u.DeputadoApi.Get(
		api_deputado.NewApiDeputadoIntputDTO(input.id, input.timeout),
	)
	if apiDeputadoOutputDTO.Error != nil {
		log.Println(apiDeputadoOutputDTO.Error)
		return NewGetDeputadoFromApiOutputDTO(nil, elapsedTime.Elapsed(), apiDeputadoOutputDTO.Error)
	}
	deputadoModel := apiDeputadoOutputDTO.Deputado

	result.Id = deputadoModel.Id
	result.Nome = deputadoModel.UltimoStatus.Nome
	result.UF = deputadoModel.UltimoStatus.SiglaUf

	// result.Exercicio
	result.Partido = deputadoModel.UltimoStatus.SiglaPartido
	// result.Legislaturas
	result.NomeCivil = deputadoModel.NomeCivil
	// result.Email = deputadoModel.UltimoStatus.Gabinete.Email
	result.Uri = deputadoModel.Uri
	// result.Url
	// result.Foto = deputadoModel.UltimoStatus.UrlFoto
	// result.Partidos
	// result.Telefone = deputadoModel.UltimoStatus.Gabinete.Telefone
	// result.Endereco
	result.DataNasc = deputadoModel.DataNascimento
	result.Naturalidade = deputadoModel.MunicipioNascimento
	// result.Atuacoes
	// result.Presencas
	// result.Cargos
	// result.Frentes

	result.UltimoStatus = deputadoModel.UltimoStatus

	result.Cpf = deputadoModel.Cpf
	result.Sexo = deputadoModel.Sexo
	result.UrlWebsite = deputadoModel.UrlWebsite
	result.RedeSocial = deputadoModel.RedeSocial
	result.DataNascimento = deputadoModel.DataNascimento
	result.DataFalecimento = deputadoModel.DataFalecimento
	result.UfNascimento = deputadoModel.UfNascimento
	result.MunicipioNascimento = deputadoModel.MunicipioNascimento
	result.Escolaridade = deputadoModel.Escolaridade

	return NewGetDeputadoFromApiOutputDTO(result, elapsedTime.Elapsed(), nil)
}
