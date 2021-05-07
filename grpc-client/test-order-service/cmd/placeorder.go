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

// placeorderCmd represents the placeorder command
var placeorderCmd = &cobra.Command{
	Use:   "placeorder",
	Short: "Place an Order",
	Long:  `Place an Order for a stock.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		exchange_id, err := cmd.Flags().GetString("exchange_id")
		symbol, err := cmd.Flags().GetString("symbol")
		qty, err := cmd.Flags().GetFloat64("qty")
		limit_price, err := cmd.Flags().GetFloat64("limit_price")
		side, err := cmd.Flags().GetString("side")
		order_type, err := cmd.Flags().GetString("order_type")
		time_in_force, err := cmd.Flags().GetString("time_in_force")

		if err != nil {
			return err
		}

		order := &orderpb.PlaceOrderReq{
			ExchangeId: exchange_id,
			Symbol:     symbol,
			Qty:        qty,
			LimitPrice: limit_price,
		}

		if side == "buy" {
			order.Side = orderpb.Side_buy
		} else if side == "sell" {
			order.Side = orderpb.Side_sell
		} else {
			order.Side = orderpb.Side_invalid_side
		}

		if order_type == "market" {
			order.OrderType = orderpb.OrderType_market
		} else if order_type == "limit" {
			order.OrderType = orderpb.OrderType_limit
		} else if order_type == "stop" {
			order.OrderType = orderpb.OrderType_stop
		} else if order_type == "stop_limit" {
			order.OrderType = orderpb.OrderType_stop_limit
		} else if order_type == "trailing_stop" {
			order.OrderType = orderpb.OrderType_trailing_stop
		} else {
			order.OrderType = orderpb.OrderType_invalid_order_type
		}

		if time_in_force == "day" {
			order.TimeInForce = orderpb.TimeInForce_day
		} else if time_in_force == "gtc" {
			order.TimeInForce = orderpb.TimeInForce_gtc
		} else if time_in_force == "opg" {
			order.TimeInForce = orderpb.TimeInForce_opg
		} else if time_in_force == "ioc" {
			order.TimeInForce = orderpb.TimeInForce_ioc
		} else if time_in_force == "fok" {
			order.TimeInForce = orderpb.TimeInForce_fok
		} else if time_in_force == "gtx" {
			order.TimeInForce = orderpb.TimeInForce_gtx
		} else if time_in_force == "gtd" {
			order.TimeInForce = orderpb.TimeInForce_gtd
		} else if time_in_force == "cls" {
			order.TimeInForce = orderpb.TimeInForce_cls
		} else {
			order.TimeInForce = orderpb.TimeInForce_invalid_time_in_force
		}

		//record starttime of request
		startTime := time.Now()

		// RPC call
		res, err := client.PlaceOrder(
			context.Background(),
			order,
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
	placeorderCmd.Flags().StringP("exchange_id", "e", "", "Add an exchange ID")
	placeorderCmd.Flags().StringP("symbol", "a", "", "Add a stock Symbol")
	placeorderCmd.Flags().Float64P("qty", "q", 0, "Add Quantity")
	placeorderCmd.Flags().StringP("side", "s", "", "Add whether to buy or sell")
	placeorderCmd.Flags().StringP("order_type", "t", "", "Add order type")
	placeorderCmd.Flags().StringP("time_in_force", "f", "", "Enter time in force. Options: day,gtc,opq,ioc,fok,gtx,gtd,cls")
	placeorderCmd.Flags().Float64P("limit_price", "l", 0.0, "Add Limit Price")
	rootCmd.AddCommand(placeorderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// placeorderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// placeorderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
