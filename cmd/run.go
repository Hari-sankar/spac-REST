package cmd

import (
	"spac-REST/api"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the  api Server",
	Long:  "Starts the api server for comments",
	Run: func(cmd *cobra.Command, args []string) {
		api.Run()
	},
}

func init() {

	rootCmd.AddCommand(startCmd)
}
