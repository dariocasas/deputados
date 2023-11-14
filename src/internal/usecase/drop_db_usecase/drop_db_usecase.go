package drop_db_usecase

import (
	"log"

	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado"
	"github.com/dariocasas/deputados3/src/internal/infra/mongodb/repo_deputado_index"
)

type DropDbUseCase struct {
	deputadoIndexRepository repo_deputado_index.DeputadoIndexRepositoryInterface
	deputadoRepository      repo_deputado.DeputadoRepositoryInterface
}

func NewDropDbUseCase(
	deputadoIndexRepository repo_deputado_index.DeputadoIndexRepositoryInterface,
	deputadoRepository repo_deputado.DeputadoRepositoryInterface,
) *DropDbUseCase {

	log.Println("NewDropDbUseCase")

	return &DropDbUseCase{
		deputadoIndexRepository: deputadoIndexRepository,
		deputadoRepository:      deputadoRepository,
	}
}

func (d DropDbUseCase) Execute() *DropDbUseCaseOutputDTO {

	log.Println("DropDbUseCase.Execute()")

	if d.deputadoRepository.CollectionExists() {
		err := d.deputadoRepository.DropCollention()
		if err != nil {
			log.Printf("DropDbUseCase error %s", err.Error())
			NewDropDbUseCaseOutputDTO(err)
		}
	}

	if d.deputadoIndexRepository.CollectionExists() {
		err := d.deputadoIndexRepository.DropCollention()
		if err != nil {
			log.Printf("DropDbUseCase error %s", err.Error())
			NewDropDbUseCaseOutputDTO(err)
		}
	}

	return NewDropDbUseCaseOutputDTO(nil)
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
