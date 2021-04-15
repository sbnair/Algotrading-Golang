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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Exchange",
	Long:  `Create a new exchange on the server through gRPC.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the data from our flags
		selected_exchange, err := cmd.Flags().GetString("selected_exchange")
		exchange_name, err := cmd.Flags().GetString("exchange_name")
		exchange_type, err := cmd.Flags().GetString("exchange_type")
		user_id, err := cmd.Flags().GetString("user_id")
		api_key, err := cmd.Flags().GetString("api_key")
		api_secret, err := cmd.Flags().GetString("api_secret")
		if err != nil {
			return err
		}
		// Create a exchange protobuffer message
		exchange := &exchangepb.Exchange{
			SelectedExchange: selected_exchange,
			ExchangeName:     exchange_name,
			ExchangeType:     exchange_type,
			UserId:           user_id,
			ApiKey:           api_key,
			ApiSecret:        api_secret,
		}
		// RPC call
		res, err := client.CreateExchange(
			context.TODO(),
			// wrap the blog message in a CreateBlog request message
			&exchangepb.CreateExchangeReq{
				Exchange: exchange,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Exchange created: %s\n", res.Exchange.Id)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("selected_exchange", "e", "", "Add an exchange")
	createCmd.Flags().StringP("exchange_name", "n", "", "Display Name of the exchange")
	createCmd.Flags().StringP("exchange_type", "t", "", "Exchange Type")
	createCmd.Flags().StringP("user_id", "u", "", "User Id")
	createCmd.Flags().StringP("api_key", "k", "", "The api key for the exchange")
	createCmd.Flags().StringP("api_secret", "s", "", "The api secret for the exchange")
	createCmd.MarkFlagRequired("selected_exchange")
	createCmd.MarkFlagRequired("exchange_name")
	createCmd.MarkFlagRequired("exchange_type")
	createCmd.MarkFlagRequired("user_id")
	createCmd.MarkFlagRequired("api_key")
	createCmd.MarkFlagRequired("api_secret")
	//rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
