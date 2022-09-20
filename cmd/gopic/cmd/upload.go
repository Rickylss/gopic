/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"fmt"

	"github.com/OSTGO/gopic/cmd/gopic/run"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewUploadCmd() *cobra.Command {
	uploadFlags := &run.UploadFlags{}

	// uploadCmd represents the upload command
	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "upload pic list",
		Long:  `upload pic list`,
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			uo := uploadFlags.NewUploadOptions(cmd, args)
			outURL := uo.CmdUpload()
			fmt.Print(outURL)
		},
	}

	addUploadFlags(uploadCmd.Flags(), uploadFlags)
	uploadCmd.MarkFlagRequired("pathList")
	return uploadCmd
}

func addUploadFlags(flags *pflag.FlagSet, uploadFlags *run.UploadFlags) {
	flags.StringVarP(&uploadFlags.Path, "path", "p", "", "")

	flags.BoolVarP(&uploadFlags.AllStorage, "all", "a", false, "")
	flags.StringSliceVarP(&uploadFlags.StorageList, "storage", "s", nil, "")
	flags.StringVarP(&uploadFlags.OutFormat, "format", "f", "", "")
	flags.BoolVarP(&uploadFlags.NameReserve, "name", "n", false, "")
}
