/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"stocker_bot/bot"
	"strconv"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the telegram bot",
	Long:  `Starts the telegram bot`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

		authorizedUserID := os.Getenv("AUTHORIZED_USER_ID")
		userID, err := strconv.ParseInt(authorizedUserID, 10, 64)
		if err != nil {
			log.Fatal("Error converting authorized user id: ", err)
		}

		config := bot.NewConfig(
			userID,
			os.Getenv("TELEGRAM_APITOKEN"),
			os.Getenv("FINNHUB_API_KEY"),
		)

		bot.Start(*config)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
