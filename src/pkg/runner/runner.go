package runner

import (
	"log"
	"sync"
)

type Runner struct {
	done chan struct{}

	concurrency int
}

func NewRunner() RunnerIterface {
	return &Runner{}
}

func (r *Runner) SetConcurrency(concurrency int) {
	r.concurrency = concurrency
	r.done = make(chan struct{}, concurrency)
}

func (r Runner) Done() {
	for i := 0; i < r.concurrency; i++ {
		r.done <- struct{}{}
		log.Printf("r.done <- struct{}{} %d", i)
	}
}

func (r Runner) GetDone() DoneChan {
	return r.done
}

func (r Runner) Run(sourceChan <-chan InputChanType, f FunctionType) <-chan OuputChanType {

	// Stage 2
	// fan out
	spreadChans := r.fanOut(sourceChan)

	// Stage 3
	// run workers
	outChannels := r.runWorkers(spreadChans, f)

	// Stage 4
	// fan in
	sinkChan := r.merge(outChannels...)

	return sinkChan
}

func (r Runner) fanOut(in <-chan InputChanType) []<-chan InputChanType {

	var outC []chan InputChanType
	for i := 0; i < r.concurrency; i++ {
		outC = append(outC,
			make(chan InputChanType),
		)
	}

	// distributes chan values in r.concurrency separated channels
	go func() {

		defer func() {
			for i := 0; i < r.concurrency; i++ {
				close(outC[i])
			}
		}()

		i := 0
		for v := range in {
			select {
			case outC[i] <- v:
				i++
				if i == r.concurrency {
					i = 0
				}
			case <-r.done:
				return
			}
		}
	}()

	// copy out channels to read-only channels
	resultChan := make([]<-chan InputChanType, r.concurrency)
	for n, ch := range outC {
		resultChan[n] = ch
	}
	return resultChan
}

func (r Runner) runWorkers(inC []<-chan InputChanType, f FunctionType) []<-chan OuputChanType {

	var outC []chan OuputChanType
	for i := 0; i < r.concurrency; i++ {
		outC = append(outC,
			make(chan OuputChanType),
		)
	}

	for i := 0; i < r.concurrency; i++ {

		// starts a worker for each concurrent channel
		go func(i int, c <-chan InputChanType, out *chan OuputChanType) {
			defer func() {
				close(*out)
			}()

			for n := range c {
				select {
				case <-r.done:
					return
				default:
					// call function
					s := f(n)
					*out <- s
				}
			}
		}(i, inC[i], &outC[i])
	}

	// copy out channels to read-only channels
	resultChan := make([]<-chan OuputChanType, r.concurrency)
	for n, ch := range outC {
		resultChan[n] = ch
	}

	return resultChan

}

func (r Runner) merge(cs ...<-chan OuputChanType) <-chan OuputChanType {
	var wg sync.WaitGroup
	out := make(chan OuputChanType)

	wg.Add(len(cs))

	output := func(c <-chan OuputChanType) {
		i := 0
		defer func() {
			wg.Done()
		}()
		for n := range c {
			select {
			case out <- n:
				i++
			case <-r.done:
				return
			}
		}
	}

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
