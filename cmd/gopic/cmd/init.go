/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"fmt"

	"github.com/OSTGO/gopic/cmd/gopic/run"
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init env",
		Long:  `init env,please use root privilege`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(run.CmdInit())
		},
	}

	return initCmd
}
