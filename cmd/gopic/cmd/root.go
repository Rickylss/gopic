/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gopic",
		Short: "A brief description of your application",
		Long:  `gopic is an opensource picPool tool`,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	// hide the help command.
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	rootCmd.AddCommand(NewConfCmd())
	rootCmd.AddCommand(NewConvertCmd())
	rootCmd.AddCommand(NewInitCmd())
	rootCmd.AddCommand(NewUpdateCmd())
	rootCmd.AddCommand(NewUploadCmd())

	return rootCmd
}
