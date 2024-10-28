package scheduler

import (
	"sfw/pkg/utils/scheduler"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var (
	Schdeduler *scheduler.Scheduler
)

func Init() {
	Schdeduler = scheduler.NewScheduler(scheduler.SchedulerOption{
		LogPrefix: "Scheduler: ",
		LogFunc:   hlog.Info,
		Debug:     true,
		Silent:    false,
		NShard:    64,
	})
}
