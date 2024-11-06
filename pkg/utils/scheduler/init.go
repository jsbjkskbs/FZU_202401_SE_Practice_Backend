package scheduler

import (
	"sfw/pkg/utils/logger"
)

var Schdeduler *Scheduler

func Init() {
	Schdeduler = NewScheduler(SchedulerOption{
		LogPrefix: "Scheduler: ",
		LogFunc:   logger.SynchronizeLogger.Info,
		Debug:     true,
		Silent:    false,
		NShard:    64,
	})
}
