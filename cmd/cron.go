package cmd

import (
	"app-api/crontab"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Cron Jobs Server",
	Run: func(cmd *cobra.Command, args []string) {
		c := cron.New(cron.WithLogger(cron.DefaultLogger))
		crontab.CrontabSchedule(c)
		c.Run()
	},
}
