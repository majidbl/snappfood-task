/*
Copyright © 2023 majidzarephysics@gmail.com. All rights reserved.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlStorage "task/storage/mysql"
	"task/util"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "run migration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migration called")
	},
}

// migrationUpCmd represents the migration command
var migrationUpCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migration up called")
		config, err := util.LoadConfig(".")
		if err != nil {
			panic(err.Error())
		}

		mysqlDB, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		if err := mysqlStorage.MigrateUp(mysqlDB); err != nil {
			panic(err.Error())
		}
	},
}

// migrationDownCmd represents the migration command
var migrationDownCmd = &cobra.Command{
	Use:   "down",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migration down called")
		config, err := util.LoadConfig(".")
		if err != nil {
			panic(err.Error())
		}

		mysqlDB, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		if err := mysqlStorage.MigrateDown(mysqlDB); err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
	migrationCmd.AddCommand(migrationUpCmd)
	migrationCmd.AddCommand(migrationDownCmd)
}
