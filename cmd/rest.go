/*
Copyright © 2023 majidzarephysics@gmail.com. All rights reserved.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"task/app"
	"task/config"
)

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "delay notification rest server",
	Long: `delay notification rest server for report order delay, assign delayed order to an free agent and report
            vendor delay`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rest called")
		c, err := config.LoadConfig(".")
		if err != nil {
			panic(err.Error())
		}
		app.NewServer(c)
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
