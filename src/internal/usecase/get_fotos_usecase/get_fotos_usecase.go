package get_fotos_usecase

import (
	"log"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado"
)

type FotosUseCase struct {
	repository repo_deputado.DeputadoRepositoryInterface
}

func NewFotosUseCase(
	deputadoRepository repo_deputado.DeputadoRepositoryInterface,
) *FotosUseCase {

	log.Println("NewFotosUseCase")

	return &FotosUseCase{

		repository: deputadoRepository,
	}
}

func (p FotosUseCase) Execute(input *FotosInputDTO) *FotosOutputDTO {

	cursor, err := p.repository.GetCursor()
	if err != nil {
		return &FotosOutputDTO{Error: err}
	}
	defer p.repository.CloseCursor(cursor)

	chanOut := make(chan *deputado.Deputado)

	go p.repository.ReadCursor(cursor, chanOut)

	for d := range chanOut {
		input.OnNewItem(
			FotoDTO{
				Id:         d.Id,
				Nome:       d.Nome,
				Partido:    d.Partido,
				UF:         d.UF,
				Url:        d.Url,
				FotoUrl:    d.UltimoStatus.UrlFoto,
				FotoUrl2:   d.UltimoStatus.UrlFoto2,
				RedeSocial: d.RedeSocial,
			},
		)
	}

	return &FotosOutputDTO{}
}
