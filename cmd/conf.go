/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"github.com/OSTGO/gopic/cmdIementaion"

	"github.com/spf13/cobra"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "config env",
	Long:  `config env`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmdIementaion.CmdConf()
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
}
