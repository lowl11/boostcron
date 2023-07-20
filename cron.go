package boostcron

import (
	"github.com/lowl11/boostcron/pkg/interfaces"
	"github.com/lowl11/boostcron/pkg/types"
	"sync"
)

type Config struct {
	ErrorHandler types.ErrorHandler
}

type Cron struct {
	errorHandler      types.ErrorHandler
	schedulersChannel chan interfaces.Scheduler
	schedulers        []interfaces.Scheduler

	mutex sync.Mutex
}

func New(config Config) *Cron {
	cron := &Cron{
		schedulersChannel: make(chan interfaces.Scheduler),
		schedulers:        make([]interfaces.Scheduler, 0),
		errorHandler:      config.ErrorHandler,
	}

	go cron.readActions()

	return cron
}

func (cron *Cron) addScheduler(scheduler interfaces.Scheduler) {
	cron.mutex.Lock()
	defer cron.mutex.Unlock()

	cron.schedulers = append(cron.schedulers, scheduler)
}

func (cron *Cron) readActions() {
	for scheduler := range cron.schedulersChannel {
		cron.addScheduler(scheduler)
	}
}
