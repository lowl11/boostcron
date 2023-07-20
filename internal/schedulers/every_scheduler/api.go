package every_scheduler

import (
	"github.com/lowl11/boostcron/pkg/interfaces"
	"github.com/lowl11/boostcron/pkg/types"
	"time"
)

func (scheduler *Scheduler) Milliseconds() interfaces.EveryScheduler {
	scheduler.duration = time.Millisecond
	return scheduler
}

func (scheduler *Scheduler) Seconds() interfaces.EveryScheduler {
	scheduler.duration = time.Second
	return scheduler
}

func (scheduler *Scheduler) Minutes() interfaces.EveryScheduler {
	scheduler.duration = time.Minute
	return scheduler
}

func (scheduler *Scheduler) Hours() interfaces.EveryScheduler {
	scheduler.duration = time.Hour
	return scheduler
}

func (scheduler *Scheduler) Do(jobAction types.CronHandler) {
	scheduler.jobAction = jobAction
	scheduler.schedulersChannel <- scheduler
}

func (scheduler *Scheduler) Action() types.CronHandler {
	return scheduler.jobAction
}

func (scheduler *Scheduler) GetDuration() time.Duration {
	return scheduler.getDuration()
}
