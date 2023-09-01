/*
Copyright © 2023 majidzarephysics@gmail.com. All rights reserved.
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"task/models"
	"task/util"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("seed called")
		config, err := util.LoadConfig(".")
		if err != nil {
			panic(err.Error())
		}

		mysqlDB, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		err = seedVendor(mysqlDB)
		if err != nil {
			log.Println(err.Error())
		}

		err = seedOrder(mysqlDB)
		if err != nil {
			log.Println(err.Error())
		}

		err = seedTrip(mysqlDB)
		if err != nil {
			log.Println(err.Error())
		}

		err = seedAgent(mysqlDB)
		if err != nil {
			log.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func seedVendor(db *gorm.DB) error {
	sampleVendors := []models.Vendor{
		{
			Name: "vendor1",
		},
		{
			Name: "vendor2",
		},
		{
			Name: "vendor3",
		},
	}

	for _, vendor := range sampleVendors {
		err := db.Create(&vendor).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedOrder(db *gorm.DB) error {
	sampleOrders := []models.Order{
		{
			VendorId:     1,
			DeliveryTime: 60,
		},
		{
			VendorId:     2,
			DeliveryTime: 35,
		},
		{
			VendorId:     3,
			DeliveryTime: 105,
		},
	}

	for _, order := range sampleOrders {
		err := db.Create(&order).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedTrip(db *gorm.DB) error {
	sampleTrips := []models.Trip{
		{
			OrderID: 1,
			Status:  models.PICKED,
		},
	}

	for _, trip := range sampleTrips {
		err := db.Create(&trip).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedAgent(db *gorm.DB) error {
	sampleAgents := []models.Agent{
		{
			Name: "Agent1",
		},
		{
			Name: "Agent2",
		},
	}

	for _, agent := range sampleAgents {
		err := db.Create(&agent).Error
		if err != nil {
			return err
		}
	}

	return nil
}
