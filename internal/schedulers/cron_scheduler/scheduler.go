package cron_scheduler

import (
	"github.com/lowl11/boostcron/pkg/interfaces"
	"github.com/lowl11/boostcron/pkg/types"
)

type Scheduler struct {
	schedulersChannel chan interfaces.Scheduler
	cronExpression    string
	fromStart         bool

	jobAction types.CronHandler
}

func New(schedulersChannel chan interfaces.Scheduler, cronExpression string) *Scheduler {
	return &Scheduler{
		schedulersChannel: schedulersChannel,
		cronExpression:    cronExpression,
	}
}
