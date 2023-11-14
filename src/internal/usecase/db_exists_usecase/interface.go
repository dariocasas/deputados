package db_exists_usecase

type DbStatusUseCaseInterface interface {
	Execute() *DbStatusUseCaseOutputDTO
}

type DbStatusUseCaseOutputDTO = struct {
	DbStatusEnum int32
	IndexCount   int32
	Error        error
}

func NewDbStatusUseCaseOutputDTO(
	dbStatusEnum int32,
	indexCount int32,
	error error,
) *DbStatusUseCaseOutputDTO {
	return &DbStatusUseCaseOutputDTO{
		DbStatusEnum: dbStatusEnum,
		IndexCount:   indexCount,
		Error:        error,
	}
}
