package deputado

// ------------------------------ Frentes ------------------------------

type Frentes []Frente

type Frente struct {
	Local   string `json:"local"`
	Periodo string `json:"periodo"`
}

func NewFrentes() *Frentes {
	return &Frentes{}
}
