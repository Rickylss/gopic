/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"fmt"

	cmdi "github.com/OSTGO/gopic/cmdIementaion"
	"github.com/spf13/cobra"
)

var uploadFlags cmdi.UploadFlags

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload pic list",
	Long:  `upload pic list`,
	Run: func(cmd *cobra.Command, args []string) {
		uo := uploadFlags.NewUploadOptions(cmd, args)
		outURL := uo.CmdUpload()
		fmt.Print(outURL)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&uploadFlags.Path, "path", "p", "", "")
	uploadCmd.Flags().BoolVarP(&uploadFlags.AllStorage, "all", "a", false, "")
	uploadCmd.Flags().StringSliceVarP(&uploadFlags.StorageList, "storage", "s", nil, "")
	uploadCmd.Flags().StringVarP(&uploadFlags.OutFormat, "format", "f", "", "")
	uploadCmd.Flags().BoolVarP(&uploadFlags.NameReserve, "name", "n", false, "")
	uploadCmd.MarkFlagRequired("pathList")
}
