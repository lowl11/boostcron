package runner

import (
	"github.com/lowl11/boostcron/pkg/interfaces"
	"github.com/lowl11/boostcron/pkg/types"
)

type Runner struct {
	scheduler    interfaces.Scheduler
	errorHandler types.ErrorHandler
}

func New(scheduler interfaces.Scheduler) *Runner {
	return &Runner{
		scheduler: scheduler,
	}
}
