package deputado

// ------------------------------ Atuacoes ------------------------------

type Cargos []Cargo

type Cargo struct {
	Cargo   string `json:"cargo"`
	Local   string `json:"local"`
	Periodo string `json:"periodo"`
}

func NewCargos() *Cargos {
	return &Cargos{}
}
