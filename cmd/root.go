package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root command",
}

func init() {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(asynqCmd)
	rootCmd.AddCommand(cronCmd)
	rootCmd.AddCommand(testCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
