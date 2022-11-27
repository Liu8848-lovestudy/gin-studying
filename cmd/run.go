package cmd

import (
	"gin-studying/controller"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:  "run",
	Long: "启动lirary小程序",
	Run: func(cmd *cobra.Command, args []string) {
		controller.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
