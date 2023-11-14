package elapsed_time

import "time"

type ElapsedTimeInterface interface {
	Elapsed() time.Duration
	Reset()
}

type ElapsedTime struct {
	start time.Time
}

func NewElapsedTime() *ElapsedTime {
	return &ElapsedTime{start: time.Now()}
}

func (t *ElapsedTime) Elapsed() time.Duration {
	return time.Since(t.start)
}

func (t *ElapsedTime) Reset() {
	t.start = time.Now()
}
