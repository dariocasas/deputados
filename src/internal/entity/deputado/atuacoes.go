package deputado

import "fmt"

// ------------------------------ Atuacoes ------------------------------

type Atuacoes struct {
	// Items     []*Atuacao
	PropostasAutoria   int `json:"propostas-sua-autoria"`
	PropostasRelatadas int `json:"propostas-relatadas"`
	Votacoes           int `json:"votacoes"`
	Discursos          int `json:"discursos"`
}

func NewAtuacoes() *Atuacoes {
	return &Atuacoes{}
}

// func (a Atuacoes) String() string {
// 	var res = "Atuacoes [\n"
// 	for _, item := range a.Items {
// 		res += fmt.Sprintf("%s\n", item)
// 	}
// 	res += "]"
// 	return res
// }

// ------------------------------ Atuacao -------------------------------

type Atuacao struct {
	Title string
	Items []*AtuacaoItem
}

func NewAtuacao() *Atuacao {
	return &Atuacao{}
}

func (a *Atuacao) NewItem() *AtuacaoItem {
	item := AtuacaoItem{}
	a.Items = append(a.Items, &item)
	return &item
}

func (a Atuacao) String() string {
	var res = fmt.Sprintf("%s [", a.Title)
	for i, item := range a.Items {
		res += item.String()
		if i < len(a.Items)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

// ------------------------------ AtuacaoItem ---------------------------

type AtuacaoItem struct {
	Tipo       string
	Quantidade int
}

func (a AtuacaoItem) String() string {
	return fmt.Sprintf("%s: %d", a.Tipo, a.Quantidade)
}
