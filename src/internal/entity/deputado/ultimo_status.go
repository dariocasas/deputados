package deputado

type UltimoStatus struct {
	Id                int       `json:"id"`
	Uri               string    `json:"uri"`
	Nome              string    `json:"nome"`
	SiglaPartido      string    `json:"siglaPartido"`
	UriPartido        string    `json:"uriPartido"`
	SiglaUf           string    `json:"siglaUf"`
	IdLegislatura     int       `json:"idLegislatura"`
	UrlFoto           string    `json:"urlFoto,omitempty"`
	UrlFoto2          string    `json:"urlFoto2,omitempty"`
	Email             string    `json:"email,omitempty"`
	Data              string    `json:"data,omitempty"`
	NomeEleitoral     string    `json:"nomeEleitoral,omitempty"`
	Gabinete          *Gabinete `json:"gabinete"`
	Situacao          string    `json:"situacao"`
	CondicaoEleitoral string    `json:"condicaoEleitoral"`
	DescricaoStatus   string    `json:"descricaoStatus"`
}

func NewUltimoStatus() *UltimoStatus {
	return &UltimoStatus{
		Gabinete: NewGabinete(),
	}
}
