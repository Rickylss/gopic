/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update gopic",
		Long:  `update from github`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("update called")
		},
	}

	return updateCmd
}
