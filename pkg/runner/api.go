package runner

import (
	"github.com/lowl11/boostcron/internal/logger"
	"github.com/lowl11/boostcron/internal/panicer"
	"github.com/lowl11/boostcron/pkg/types"
	"time"
)

func (runner *Runner) ErrorHandler(handler types.ErrorHandler) *Runner {
	runner.errorHandler = handler
	return runner
}

func (runner *Runner) StartTicker() {
	if runner.fromStart {
		runner.runAction()
	}

	for {
		ticker := time.NewTicker(runner.scheduler.GetDuration())

		<-ticker.C

		runner.runAction()
		ticker.Reset(runner.scheduler.GetDuration())
	}
}

func (runner *Runner) FromStart(fromStart bool) *Runner {
	runner.fromStart = fromStart
	return runner
}

func (runner *Runner) runAction() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(panicer.Handle(err), "PANIC RECOVERED")
		}
	}()

	jobAction := runner.scheduler.Action()

	if err := jobAction(); err != nil {
		if runner.errorHandler != nil {
			if innerError := runner.errorHandler(err); innerError != nil {
				logger.Error(innerError, "Cron error handler error")
			}
		}

		logger.Error(err, "Execute job action error")
	}
}
