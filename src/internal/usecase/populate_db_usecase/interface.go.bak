package populate_db_usecase

import (
	"time"

	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
)

type PopulateDbUseCaseInterface interface {
	Execute(input PopulateDbInputDTO) *PopulateDbOutputDTO
}

type PopulateDbInputDTO = struct {
	TimeOut   int
	OnNewItem OnNewItem
}

type PopulateDbOutputDTO = struct {
	ItemsProcessed int
	Duration       time.Duration
	Error          error
}

type PartialTime struct {
	GetApi   time.Duration
	DbRead   time.Duration
	GetPage  time.Duration
	ScanPage time.Duration
	DbWrite  time.Duration
}

type OnNewItem func(
	id int,
	name string,
	error error,
	elapsedTime elapsed_time.ElapsedTimeInterface,
	partialTime PartialTime,
)
