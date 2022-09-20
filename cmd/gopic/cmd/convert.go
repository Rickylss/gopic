/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/OSTGO/gopic/cmd/gopic/run"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewConvertCmd() *cobra.Command {
	covertFlags := &run.ConvertFlags{}
	convertCmd := &cobra.Command{
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

	addConvertFlags(convertCmd.Flags(), covertFlags)
	convertCmd.MarkFlagRequired("dir")

	return convertCmd
}

func addConvertFlags(flags *pflag.FlagSet, covertFlags *run.ConvertFlags) {
	flags.StringVarP(&covertFlags.CovertPath, "covertPath", "c", "./", "")
	flags.StringVarP(&covertFlags.OutDir, "dir", "d", "", "")
	flags.BoolVarP(&covertFlags.Recurse, "recurse", "r", false, "")

	flags.BoolVarP(&covertFlags.AllStorage, "all", "a", false, "")
	flags.StringSliceVarP(&covertFlags.StorageList, "storage", "s", nil, "")
	flags.StringVarP(&covertFlags.OutFormat, "format", "f", "", "")
	flags.BoolVarP(&covertFlags.NameReserve, "name", "n", false, "")
}
