package cmd

import (
	"gin-studying/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:  "init",
	Long: "初始化数据表",
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateTables()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
