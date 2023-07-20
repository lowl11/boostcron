package runner

import (
	"github.com/lowl11/boostcron/pkg/types"
	"github.com/lowl11/lazylog/log"
	"time"
)

func (runner *Runner) ErrorHandler(handler types.ErrorHandler) *Runner {
	runner.errorHandler = handler
	return runner
}

func (runner *Runner) StartTicker() {
	for {
		ticker := time.NewTicker(runner.scheduler.GetDuration())

		<-ticker.C

		runner.runAction()
		ticker.Reset(runner.scheduler.GetDuration())
	}

}

func (runner *Runner) runAction() {
	jobAction := runner.scheduler.Action()

	if err := jobAction(); err != nil {
		if runner.errorHandler != nil {
			if innerError := runner.errorHandler(err); innerError != nil {
				log.Error(innerError, "Cron error handler error")
			}
		}

		log.Error(err, "Execute job action error")
	}
}
