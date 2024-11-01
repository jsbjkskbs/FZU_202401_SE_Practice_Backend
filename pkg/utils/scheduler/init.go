package scheduler

import (

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var (
	Schdeduler *Scheduler
)

func Init() {
	Schdeduler = NewScheduler(SchedulerOption{
		LogPrefix: "Scheduler: ",
		LogFunc:   hlog.Info,
		Debug:     true,
		Silent:    false,
		NShard:    64,
	})
}
