package model

import (
	"log"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado"
)

type Deputado struct {
	Id                  int                    `json:"id"`
	Uri                 string                 `json:"uri,omitempty"`
	NomeCivil           string                 `json:"nomecivil"`
	UltimoStatus        *deputado.UltimoStatus `json:"ultimoStatus"`
	Cpf                 string                 `json:"cpf"`
	Sexo                string                 `json:"sexo"`
	UrlWebsite          string                 `json:"urlWebsite"`
	RedeSocial          []string               `json:"redeSocial"`
	DataNascimento      string                 `json:"dataNascimento"`
	DataFalecimento     string                 `json:"dataFalecimento,omitempty"`
	UfNascimento        string                 `json:"ufNascimento"`
	MunicipioNascimento string                 `json:"municipioNascimento"`
	Escolaridade        string                 `json:"escolaridade"`
}

func NewDeputado() (*Deputado, error) {
	log.Println("NewDeputado")
	return &Deputado{
		UltimoStatus: deputado.NewUltimoStatus(),
	}, nil
}
