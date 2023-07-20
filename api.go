package boostcron

import (
	"github.com/lowl11/boostcron/internal/schedulers/cron_scheduler"
	"github.com/lowl11/boostcron/internal/schedulers/every_scheduler"
	"github.com/lowl11/boostcron/pkg/runner"
	"time"
)

func (cron *Cron) Every(every int) *every_scheduler.Scheduler {
	return every_scheduler.New(cron.schedulersChannel, every)
}

func (cron *Cron) Cron(cronExpression string) *cron_scheduler.Scheduler {
	return cron_scheduler.New(cron.schedulersChannel, cronExpression)
}

func (cron *Cron) Run() {
	close(cron.schedulersChannel)

	time.Sleep(time.Millisecond * 250)

	for _, scheduler := range cron.schedulers {
		go runner.
			New(scheduler).
			ErrorHandler(cron.errorHandler).
			StartTicker()
	}

	infinite := make(chan bool, 1)
	<-infinite
}

func (cron *Cron) RunAsync() {
	go cron.Run()
}
