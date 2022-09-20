/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	cmdi "github.com/OSTGO/gopic/cmdIementaion"
	"github.com/spf13/cobra"
)

var covertFlags cmdi.ConvertFlags

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert pics",
	Long:  `convert pics`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		co := covertFlags.NewConvertOptions(cmd, args)
		err := co.CmdConvert()
		if err != nil {
			panic(err)
		}
		fmt.Println("convert success!")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&covertFlags.CovertPath, "covertPath", "c", "./", "")
	convertCmd.Flags().BoolVarP(&covertFlags.AllStorage, "all", "a", false, "")
	convertCmd.Flags().BoolVarP(&covertFlags.Recurse, "recurse", "r", false, "")
	convertCmd.Flags().StringSliceVarP(&covertFlags.StorageList, "storage", "s", nil, "")
	convertCmd.Flags().StringVarP(&covertFlags.OutFormat, "format", "f", "", "")
	convertCmd.Flags().StringVarP(&covertFlags.OutDir, "dir", "d", "", "")
	convertCmd.Flags().BoolVarP(&covertFlags.NameReserve, "name", "n", false, "")
	convertCmd.MarkFlagRequired("dir")
}
