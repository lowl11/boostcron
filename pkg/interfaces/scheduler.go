package interfaces

import (
	"github.com/lowl11/boostcron/pkg/types"
	"time"
)

type Scheduler interface {
	Action() types.CronHandler
	GetDuration() time.Duration
}

type BaseScheduler interface {
	Do(jobAction types.CronHandler)
}

type EveryScheduler interface {
	BaseScheduler

	Milliseconds() EveryScheduler
	Seconds() EveryScheduler
	Minutes() EveryScheduler
	Hours() EveryScheduler
}

type CronScheduler interface {
	BaseScheduler
}
