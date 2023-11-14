package deputado

import "fmt"

// ------------------------------------ Presencas -----------------------------------

type Presencas struct {
	// Items                            []*Presenca
	PresencaPlenario                 int
	PresencaComicoes                 int
	AusenciasJustificadasPlenario    int
	AusenciasJustificadasComicoes    int
	AusenciasNaoJustificadasPlenario int
	AusenciasNaoJustificadasComicoes int
}

func NewPresencas() *Presencas {
	return &Presencas{}
}

// func (a Presencas) String() string {
// 	var res = "Presencas [\n"
// 	for _, item := range a.Items {
// 		res += fmt.Sprintf("%s\n", item)
// 	}
// 	res += "]"
// 	return res
// }

// func (a *Presencas) Add(item *Presenca) {
// 	a.Items = append(a.Items, item)
// }

// ------------------------------------ Presenca ------------------------------------

type Presenca struct {
	Title string
	Items []*PresencaItem
}

func NewPresenca() *Presenca {
	return &Presenca{}
}

func (a *Presenca) Newitem() *PresencaItem {
	item := PresencaItem{}
	a.Items = append(a.Items, &item)
	return &item
}

func (a Presenca) String() string {
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

// ------------------------------------ PresencaItem --------------------------------

type PresencaItem struct {
	Tipo       string
	Quantidade int
}

func (a PresencaItem) String() string {
	return fmt.Sprintf("%s: %d", a.Tipo, a.Quantidade)
}
