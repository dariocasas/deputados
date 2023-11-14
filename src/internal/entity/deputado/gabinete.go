package deputado

type Gabinete struct {
	Nome     string `json:"nome"`
	Predio   string `json:"predio"`
	Sala     string `json:"sala"`
	Andar    string `json:"andar"`
	Telefone string `json:"telefone"`
	Email    string `json:"email,omitempty"`
}

func NewGabinete() *Gabinete {
	return &Gabinete{}
}
