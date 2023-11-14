package scan_page

import (
	"time"

	deputado "github.com/dariocasas/deputados3/src/internal/entity/deputado"
)

type ScanPageDeputadoInterface interface {
	Parse(ScanPageDeputadoInputDTO) *ScanPageDeputadoOutputDTO
}

// --------------------- INPUT DTO ---------------------

type ScanPageDeputadoInputDTO struct {
	id   int
	Html string
}

func NewScanPageDeputadoInputDTO(id int, pageHtml string) ScanPageDeputadoInputDTO {
	return ScanPageDeputadoInputDTO{
		id:   id,
		Html: pageHtml,
	}
}

// --------------------- OUTPUT DTO --------------------

type ScanPageDeputadoOutputDTO struct {
	Deputado *deputado.Deputado
	Duration time.Duration
	Error    error
}

func NewScanPageDeputadoOutputDTO(deputado *deputado.Deputado, duration time.Duration, error error) *ScanPageDeputadoOutputDTO {
	return &ScanPageDeputadoOutputDTO{
		Deputado: deputado,
		Duration: duration,
		Error:    error,
	}
}
