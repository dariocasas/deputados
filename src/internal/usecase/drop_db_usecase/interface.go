package drop_db_usecase

type DropDbUseCaseOutputDTO = struct {
	Error error
}

func NewDropDbUseCaseOutputDTO(
	error error,
) *DropDbUseCaseOutputDTO {
	return &DropDbUseCaseOutputDTO{
		Error: error,
	}
}

type DropDbUseCaseInterface interface {
	Execute() *DropDbUseCaseOutputDTO
}
