package interfaces

import (
	"github.com/lowl11/boostcron/pkg/types"
	"time"
)

type Scheduler interface {
	Action() types.CronHandler
	GetDuration() time.Duration
}
