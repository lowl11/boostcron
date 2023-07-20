package every_scheduler

import (
	"github.com/lowl11/boostcron/pkg/interfaces"
	"github.com/lowl11/boostcron/pkg/types"
	"time"
)

type Scheduler struct {
	schedulersChannel chan interfaces.Scheduler
	every             int

	duration time.Duration

	jobAction types.CronHandler
}

func New(schedulersChannel chan interfaces.Scheduler, every int) *Scheduler {
	return &Scheduler{
		schedulersChannel: schedulersChannel,
		every:             every,
	}
}
