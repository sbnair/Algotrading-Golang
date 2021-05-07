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
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	orderpb "github.com/vikjdk7/Algotrading-Golang/order-service/proto"
)

// cancelorderCmd represents the cancelorder command
var cancelorderCmd = &cobra.Command{
	Use:   "cancelorder",
	Short: "Cancel an order",
	Long:  `Cancel an order by order id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		exchange_id, err := cmd.Flags().GetString("exchange_id")
		order_id, err := cmd.Flags().GetString("order_id")
		if err != nil {
			return err
		}

		cancelOrder := &orderpb.CancelOrderReq{
			ExchangeId: exchange_id,
			OrderId:    order_id,
		}
		//record starttime of request
		startTime := time.Now()

		res, err := client.CancelOrder(
			context.Background(),
			cancelOrder,
		)
		if err != nil {
			return err
		}
		//fmt.Println(res)
		jsonBytes, _ := json.MarshalIndent(res, "", "    ")
		fmt.Println(string(jsonBytes))
		//fmt.Println(res)

		diff := time.Since(startTime)
		fmt.Print("Time taken for the operation: ")
		fmt.Println(diff)
		return nil
	},
}

func init() {
	cancelorderCmd.Flags().StringP("exchange_id", "e", "", "Add an exchange ID")
	cancelorderCmd.Flags().StringP("order_id", "o", "", "Add an Order Id")
	rootCmd.AddCommand(cancelorderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelorderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelorderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
