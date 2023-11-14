package model

import (
	"log"

	"github.com/dariocasas/deputados3/src/internal/entity/deputado_index"
)

type Deputado struct {
	Id            int    `json:"id" bson:"id"`
	Uri           string `json:"uri,omitempty"`
	Nome          string `json:"nome"`
	SiglaPartido  string `json:"siglaPartido"`
	UriPartido    string `json:"uriPartido"`
	SiglaUf       string `json:"siglaUf"`
	IdLegislatura int    `json:"idLegislatura"`
	UrlFoto       string `json:"urlFoto,omitempty"`
	Email         string `json:"email,omitempty"`
}

type Deputados []*Deputado

func NewDeputados(ids []*Deputado) (Deputados, error) {
	log.Println("NewDeputados(ids...)")
	return ids, nil
}

func (d Deputados) Ids() deputado_index.DeputadoIndex {
	var res deputado_index.DeputadoIndex
	for _, v := range d {
		deputadoId := deputado_index.DeputadoId{Id: v.Id}
		res = append(res, &deputadoId)
	}
	return res
}
