package runner

type InputChanType struct {
	Value any
}

type OuputChanType struct {
	Value any
}

type FunctionType = func(InputChanType) OuputChanType
type DoneChan = chan struct{}

type RunnerIterface interface {
	Run(<-chan InputChanType, FunctionType) <-chan OuputChanType
	Done()
	GetDone() DoneChan
	SetConcurrency(concurrency int)
}
