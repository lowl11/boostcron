package cron_scheduler

import (
	"github.com/lowl11/boostcron/pkg/types"
	"time"
)

func (scheduler *Scheduler) Action() types.CronHandler {
	return scheduler.jobAction
}

func (scheduler *Scheduler) GetDuration() time.Duration {
	return scheduler.getDuration(scheduler.cronExpression)
}

func (scheduler *Scheduler) Do(jobAction types.CronHandler) {
	scheduler.jobAction = jobAction
	scheduler.schedulersChannel <- scheduler
}
