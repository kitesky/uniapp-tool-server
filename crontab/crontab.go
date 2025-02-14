package crontab

import (
	"app-api/utils"

	"github.com/robfig/cron/v3"
)

func CrontabSchedule(cron *cron.Cron) {
	cron.AddFunc("@every 5s", func() {
		utils.ZapLog().Info("cron", "Every 5s")
	})
}
