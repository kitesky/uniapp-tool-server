package cmd

import (
	"log"

	"app-api/boot"
	"app-api/jobs"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

var asynqCmd = &cobra.Command{
	Use:   "asynq",
	Short: "Asynq Task Server",
	Run: func(cmd *cobra.Command, args []string) {
		mux := asynq.NewServeMux()
		jobs.JobSchedule(mux)

		if err := boot.Asynq.Server.Run(mux); err != nil {
			log.Fatal(err)
		}
	},
}
