package scan_page

import (
	"log"
	"strconv"
	"strings"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
	"github.com/dariocasas/deputados3/src/pkg/class_finder"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
	"golang.org/x/net/html"
)

type ScanPageDeputado struct {
	deputado    *deputado.Deputado
	classFinder class_finder.ClassFinderInterface
}

func NewScanPageDeputado(classFinder class_finder.ClassFinderInterface) *ScanPageDeputado {
	return &ScanPageDeputado{
		classFinder: classFinder,
	}
}

func (s *ScanPageDeputado) Parse(scanPageDeputadoInputDTO ScanPageDeputadoInputDTO) *ScanPageDeputadoOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	s.deputado = deputado.NewDeputado(scanPageDeputadoInputDTO.id)

	tokenizer := html.NewTokenizer(strings.NewReader(scanPageDeputadoInputDTO.Html))

	s.scanInformacoesDeputado(tokenizer)
	s.scanAtuacaoDeputado(tokenizer)
	s.scanPresencaDeputado(tokenizer)
	s.scanCargosDeputado(tokenizer)
	s.scanFrentesDeputado(tokenizer)

	return NewScanPageDeputadoOutputDTO(s.deputado, elapsedTime.Elapsed(), nil)
}

func (s ScanPageDeputado) scanInformacoesDeputado(tokeniker *html.Tokenizer) {

	tagFound := s.classFinder.SearchClass(tokeniker, "identificacao-deputado", "")
	if tagFound != nil {
		for _, attrib := range tagFound.Token.Attr {
			switch {
			case attrib.Key == "data-ide":
				s.deputado.Id, _ = strconv.Atoi(attrib.Val)
			case attrib.Key == "data-nome":
				s.deputado.Nome = attrib.Val
			}
		}
	}

	tagFound = s.classFinder.SearchClass(tokeniker, "foto-deputado", "")
	if tagFound != nil {
		tagFound := s.classFinder.SearchNextToken(tokeniker, "img")
		s.deputado.UltimoStatus.UrlFoto2 = tagFound.Token.Attr[0].Val
	}

	s.classFinder.SearchClass(tokeniker, "informacoes-deputado", "")
	if tagFound != nil {
		s.classFinder.ScanToNextTextToken(tokeniker,
			func(text string, index int, stop *bool) {
				switch {
				case index == 2:
					s.deputado.NomeCivil = text
				case index == 5:
					data := strings.Split(text, "-")
					if len(data) > 0 {
						s.deputado.Partido = strings.Trim(data[0], " ")
						s.deputado.UF = strings.Trim(data[1], " ")
					}
					*stop = true
				}
			},
		)
	}

	s.classFinder.SearchClass(tokeniker, "email", "")
	if tagFound != nil {
		s.classFinder.ScanToNextTextToken(tokeniker,
			func(text string, index int, stop *bool) {
				switch {
				case index == 0:
					s.deputado.UltimoStatus.Email = text // 12
				case index == 3:
					s.deputado.UltimoStatus.Gabinete.Telefone = text //15
				case index == 6:
					s.deputado.Endereco = strings.Join( //18
						strings.Fields(
							strings.ReplaceAll(text, "\n", ""),
						), " ")
				case index == 9: //21
					s.deputado.DataNasc = text
				case index == 12:
					s.deputado.Naturalidade = text
					*stop = true
				}
			},
		)
	}

	s.classFinder.ScanToEndToken(tokeniker, "ul")
}

func (s ScanPageDeputado) scanAtuacaoDeputado(tokeniker *html.Tokenizer) {

	s.classFinder.SearchClass(tokeniker, "l-cards-atuacao", "")
	for i := 0; i < 4; i++ {
		s.classFinder.ScanToNextTextToken(tokeniker,
			func(text string, index int, stop *bool) {
				s.scanCardsAtuacao(tokeniker, i, s.deputado.Atuacoes)
				*stop = true
			},
		)

	}
}

func (s ScanPageDeputado) scanCardsAtuacao(tokeniker *html.Tokenizer, item int, atuacoes *deputado.Atuacoes) {

	s.classFinder.SearchClass(tokeniker, "atuacao__quantidade", "")
	s.classFinder.ScanToNextTextToken(tokeniker,
		func(text string, index int, stop *bool) {
			val, err := strconv.Atoi(text)
			if err != nil {
				log.Print(err)
			} else {
				switch item {
				case 0:
					atuacoes.PropostasAutoria = val
				case 1:
					atuacoes.PropostasRelatadas = val
				case 2:
					atuacoes.Votacoes = val
				case 3:
					atuacoes.Discursos = val
				}
			}
			*stop = true
		})
}

// ---------------------------------------------------------------------------------

func (s ScanPageDeputado) scanPresencaDeputado(tokeniker *html.Tokenizer) {
	// Presença table 2 col 3 rows
	// skip section 3 rows and read 2 cols of 3 rows
	s.classFinder.SearchClass(tokeniker, "presencas__content", "")
	s.classFinder.SearchClass(tokeniker, "presencas__section-content", "")

	// 2 sections (cols)
	for i := 0; i < 2; i++ {
		s.scanPresenca(tokeniker, s.deputado.Presencas, i)
	}
}

func (s ScanPageDeputado) scanPresenca(tokeniker *html.Tokenizer, presencas *deputado.Presencas, item int) {

	s.classFinder.SearchClass(tokeniker, "presencas__subsection", "")

	s.classFinder.SearchClass(tokeniker, "presencas__subsection-content", "")
	for i := 0; i < 3; i++ {

		idx := (item * 3) + i

		s.classFinder.SearchClass(tokeniker, "presencas__label", "")
		s.classFinder.SearchClass(tokeniker, "presencas__qtd", "")
		s.classFinder.ScanToNextTextToken(tokeniker,
			func(text string, index int, stop *bool) {

				if len(text) > 1 {
					split := strings.Split(text, " ")
					val, err := strconv.Atoi(split[0])

					if err != nil {
						log.Print(err)
					} else {
						switch idx {
						case 0:
							presencas.PresencaPlenario = val
						case 1:
							presencas.PresencaComicoes = val
						case 2:
							presencas.AusenciasJustificadasPlenario = val
						case 3:
							presencas.AusenciasJustificadasComicoes = val
						case 4:
							presencas.AusenciasNaoJustificadasPlenario = val
						case 5:
							presencas.AusenciasNaoJustificadasComicoes = val
						}
					}
				}
				*stop = true
			})
	}
}

// ---------------------------------------------------------------------------------

func (s ScanPageDeputado) scanCargosDeputado(tokeniker *html.Tokenizer) {

	cargos := deputado.NewCargos()

	defer func() {
		s.deputado.Cargos = cargos
	}()

	// Presença table 2 col 3 rows
	// skip section 3 rows and read 2 cols of 3 rows
	for {

		found1 := s.classFinder.SearchClass(tokeniker, "cargos-deputado", "titulo-interno")
		if found1 == nil {
			break
		}

		cargo := &deputado.Cargo{}
		for {
			found2 := s.classFinder.SearchClass(tokeniker, "cargos-deputado__cargo", "cargos-deputado__agrupador")
			if found2 == nil {
				break
			}
			s.classFinder.ScanToNextTextToken(
				tokeniker,
				func(text string, index int, stop *bool) {
					cargo.Cargo = text
					*stop = true
				})

			s.classFinder.SearchClass(tokeniker, "cargos-deputado__local", "")
			s.classFinder.ScanToNextTextToken(
				tokeniker,
				func(text string, index int, stop *bool) {
					cargo.Local = text
					*stop = true
				})

			s.classFinder.SearchClass(tokeniker, "cargos-deputado__periodo", "")
			s.classFinder.ScanToNextTextToken(
				tokeniker,
				func(text string, index int, stop *bool) {
					cargo.Periodo = text
					*cargos = append(*cargos, *cargo)
					*stop = true
				})

		}

	}
}

// ---------------------------------------------------------------------------------

func (s ScanPageDeputado) scanFrentesDeputado(tokeniker *html.Tokenizer) {
	frentes := deputado.NewFrentes()

	defer func() {
		s.deputado.Frentes = frentes
	}()

	if s.classFinder.SearchClass(tokeniker, "cargos-deputado__lista", "") != nil {

		frente := &deputado.Frente{}
		for {
			found := s.classFinder.SearchClass(tokeniker, "cargos-deputado__local", "gastos-deputado")
			if found == nil {
				break
			}

			s.classFinder.ScanToNextTextToken(
				tokeniker,
				func(text string, index int, stop *bool) {
					frente.Local = text
					*stop = true
				})

			s.classFinder.SearchClass(tokeniker, "cargos-deputado__periodo", "")
			s.classFinder.ScanToNextTextToken(
				tokeniker,
				func(text string, index int, stop *bool) {
					frente.Periodo = text
					*frentes = append(*frentes, *frente)
					*stop = true
				})

		}

	}
}
