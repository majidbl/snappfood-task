/*
Copyright © 2023 majidzarephysics@gmail.com. All rights reserved.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"task/app"
	"task/util"
)

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "delay notification rest server",
	Long: `delay notification rest server for report order delay, assign delayed order to an free agent and report
            vendor delay`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rest called")
		c, err := util.LoadConfig(".")
		if err != nil {
			panic(err.Error())
		}
		api.NewServer(c)
	},
}

func init() {
	rootCmd.AddCommand(restCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
