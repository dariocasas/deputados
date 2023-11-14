package get_fotos_usecase

type FotosUsecaseInterface interface {
	Execute(input *FotosInputDTO) *FotosOutputDTO
}

type FotosInputDTO = struct {
	OnNewItem OnNewItem
}

type FotoDTO = struct {
	Id         int
	FotoUrl    string
	FotoUrl2   string
	Nome       string
	Partido    string
	UF         string
	Url        string
	RedeSocial []string
}

type FotosOutputDTO = struct {
	Error error
}

type OnNewItem func(
	foto FotoDTO,
)
