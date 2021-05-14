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
	strategypb "github.com/vikjdk7/Algotrading-Golang/strategy-service/proto"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a strategy",
	Long:  `Update a Strategy by its Id in MongoDB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the flags from CLI
		id, err := cmd.Flags().GetString("id")
		strategy_name, err := cmd.Flags().GetString("strategy_name")
		selected_exchange, err := cmd.Flags().GetString("selected_exchange")
		base_order_size, err := cmd.Flags().GetFloat64("base_order_size")
		safety_order_size, err := cmd.Flags().GetFloat64("safety_order_size")
		max_safety_trade_count, err := cmd.Flags().GetString("max_safety_trade_count")
		max_active_safety_trade_count, err := cmd.Flags().GetString("max_active_safety_trade_count")
		price_devation, err := cmd.Flags().GetString("price_devation")
		safety_order_volume_scale, err := cmd.Flags().GetString("safety_order_volume_scale")
		safety_order_step_scale, err := cmd.Flags().GetString("safety_order_step_scale")
		take_profit, err := cmd.Flags().GetString("take_profit")
		target_profit, err := cmd.Flags().GetString("target_profit")
		allocate_funds_to_strategy, err := cmd.Flags().GetString("allocate_funds_to_strategy")
		stock, err := cmd.Flags().GetStringSlice("stock")

		deal := make([]*strategypb.Stock, len(stock))
		for i, v := range stock {
			deal[i] = &strategypb.Stock{
				StockName: v,
			}
		}
		// Create an UpdateExchangeReq
		strategy := &strategypb.Strategy{
			Id:                        id,
			StrategyName:              strategy_name,
			SelectedExchange:          selected_exchange,
			BaseOrderSize:             base_order_size,
			SafetyOrderSize:           safety_order_size,
			MaxSafetyTradeCount:       max_safety_trade_count,
			MaxActiveSafetyTradeCount: max_active_safety_trade_count,
			PriceDevation:             price_devation,
			SafetyOrderVolumeScale:    safety_order_volume_scale,
			SafetyOrderStepScale:      safety_order_step_scale,
			TakeProfit:                take_profit,
			TargetProfit:              target_profit,
			AllocateFundsToStrategy:   allocate_funds_to_strategy,
			Stock:                     deal,
		}

		//record starttime of request
		startTime := time.Now()

		res, err := client.UpdateStrategy(context.Background(),
			&strategypb.UpdateStrategyReq{
				Strategy: strategy,
			})
		if err != nil {
			return err
		}

		//fmt.Println(res)
		jsonBytes, _ := json.MarshalIndent(res, "", "    ")
		fmt.Println(string(jsonBytes))

		diff := time.Since(startTime)
		fmt.Println("Time taken for the operation: ")
		fmt.Println(diff)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("id", "i", "", "Add a strategy id")
	updateCmd.Flags().StringP("strategy_name", "n", "", "Add a strategy name")
	updateCmd.Flags().StringP("selected_exchange", "e", "", "Add a selected exchange")

	updateCmd.Flags().Float64P("base_order_size", "b", 0.0, "Add base order size")
	updateCmd.Flags().Float64P("safety_order_size", "s", 0.0, "Add safety order size")

	updateCmd.Flags().StringP("max_safety_trade_count", "t", "", "The maximum number of safety orders the bot can use for one deal.")
	updateCmd.Flags().StringP("max_active_safety_trade_count", "w", "", "The maximum number of safety orders the bot can use for one deal.")
	updateCmd.Flags().StringP("price_devation", "p", "", "Add price_devation")
	updateCmd.Flags().StringP("safety_order_volume_scale", "v", "", "Add safety_order_volume_scale")
	updateCmd.Flags().StringP("safety_order_step_scale", "c", "", "Add safety_order_step_scale")
	updateCmd.Flags().StringP("take_profit", "m", "", "Add take_profit")
	updateCmd.Flags().StringP("target_profit", "z", "", "Add target_profit")
	updateCmd.Flags().StringP("allocate_funds_to_strategy", "f", "", "Add allocate_funds_to_strategy")
	updateCmd.Flags().StringSliceP("stock", "d", []string{}, "Add stock")
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
