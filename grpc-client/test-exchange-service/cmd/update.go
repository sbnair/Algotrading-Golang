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
	"context"
	"fmt"

	"github.com/spf13/cobra"
	exchangepb "github.com/vikjdk7/Algotrading-Golang/exchange-service/proto"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Find an exchange by its ID",
	Long: `Find an exchange by it's mongoDB Unique identifier.
	
	If no exchange is found for the ID it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the flags from CLI
		id, err := cmd.Flags().GetString("id")
		selected_exchange, err := cmd.Flags().GetString("selected_exchange")
		exchange_name, err := cmd.Flags().GetString("exchange_name")
		exchange_type, err := cmd.Flags().GetString("exchange_type")
		user_id, err := cmd.Flags().GetString("user_id")
		api_key, err := cmd.Flags().GetString("api_key")
		api_secret, err := cmd.Flags().GetString("api_secret")

		// Create an UpdateExchangeReq
		exchange := &exchangepb.Exchange{
			Id:               id,
			SelectedExchange: selected_exchange,
			ExchangeName:     exchange_name,
			ExchangeType:     exchange_type,
			UserId:           user_id,
			ApiKey:           api_key,
			ApiSecret:        api_secret,
		}

		res, err := client.UpdateExchange(context.Background(),
			&exchangepb.UpdateExchangeReq{
				Exchange: exchange,
			})
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("id", "i", "", "The id of the exchange")
	updateCmd.Flags().StringP("selected_exchange", "e", "", "Add an exchange")
	updateCmd.Flags().StringP("exchange_name", "n", "", "Display Name of the exchange")
	updateCmd.Flags().StringP("exchange_type", "t", "", "Exchange Type")
	updateCmd.Flags().StringP("user_id", "u", "", "User Id")
	updateCmd.Flags().StringP("api_key", "k", "", "API key for the exchange")
	updateCmd.Flags().StringP("api_secret", "s", "", "API secret for the change")
	updateCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
