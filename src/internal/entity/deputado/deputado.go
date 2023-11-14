package deputado

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deputado struct {
	Oid  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id   int                `json:"id" bson:"id"`
	Nome string             `json:"nome" bson:"nome"`
	UF   string             `json:"uf" bson:"uf"`
	// Exercicio    string             `json:"exercicio" bson:"exercicio"`
	Partido string `json:"partido" bson:"partido"`
	// Legislaturas []string `json:"legislaturas" bson:"legislaturas"`
	NomeCivil string `json:"nomecivil,omitempty" bson:"nomecivil,omitempty"`
	// Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Uri string `json:"uri,omitempty" bson:"uri,omitempty"`
	Url string `json:"url,omitempty" bson:"url,omitempty"`
	//Foto         string             `json:"foto,omitempty" bson:"foto,omitempty"`
	// Partidos     []string                  `json:"partidos,omitempty" bson:"partidos,omitempty"`
	// Telefone     string     `json:"telefone,omitempty" bson:"telefone,omitempty"`
	Endereco     string `json:"endereco,omitempty" bson:"endereco,omitempty"`
	DataNasc     string `json:"datanasc,omitempty" bson:"datanasc,omitempty"`
	Naturalidade string `json:"naturalidade,omitempty" bson:"naturalidade,omitempty"`

	Atuacoes  *Atuacoes  `json:"atuacoes,omitempty" bson:"atuacoes,omitempty"`
	Presencas *Presencas `json:"presencas,omitempty" bson:"presencas,omitempty"`
	Cargos    *Cargos    `json:"cargos,omitempty" bson:"cargos,omitempty"`
	Frentes   *Frentes   `json:"frentes,omitempty" bson:"frentes,omitempty"`

	UltimoStatus *UltimoStatus `json:"ultimoStatus"`

	Cpf                 string   `json:"cpf"`
	Sexo                string   `json:"sexo"`
	UrlWebsite          string   `json:"urlWebsite"`
	RedeSocial          []string `json:"redeSocial"`
	DataNascimento      string   `json:"dataNascimento"`
	DataFalecimento     string   `json:"dataFalecimento,omitempty"`
	UfNascimento        string   `json:"ufNascimento"`
	MunicipioNascimento string   `json:"municipioNascimento"`
	Escolaridade        string   `json:"escolaridade"`
}

func NewDeputado(id int) *Deputado {
	return &Deputado{
		Id:           id,
		Atuacoes:     NewAtuacoes(),
		Presencas:    NewPresencas(),
		Cargos:       NewCargos(),
		UltimoStatus: NewUltimoStatus(),
	}
}
