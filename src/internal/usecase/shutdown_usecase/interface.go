package shutdown_usecase

type ShutdownUseCaseInterface interface {
	Execute(input *ShutdownInputDTO) *ShutdownOutputDTO
}

type ShutdownInputDTO = struct{}

type ShutdownOutputDTO = struct {
	Error error
}
