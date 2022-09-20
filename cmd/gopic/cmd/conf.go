/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	cmdIementaion "github.com/OSTGO/gopic/cmd/gopic/run"

	"github.com/spf13/cobra"
)

func NewConfCmd() *cobra.Command {
	confCmd := &cobra.Command{
		Use:   "conf",
		Short: "config env",
		Long:  `config env`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmdIementaion.CmdConf()
		},
	}

	return confCmd
}
