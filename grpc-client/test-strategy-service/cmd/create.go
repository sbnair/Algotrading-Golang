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
	"time"

	"github.com/spf13/cobra"
	strategypb "github.com/vikjdk7/Algotrading-Golang/strategy-service/proto"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Strategy",
	Long:  `Create a new strategy on the server through gRPC.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the data from our flags
		bot_name, err := cmd.Flags().GetString("bot_name")
		selected_strategy, err := cmd.Flags().GetString("selected_strategy")
		bot_type, err := cmd.Flags().GetString("bot_type")
		pairs, err := cmd.Flags().GetString("pairs")
		strategy_type, err := cmd.Flags().GetString("strategy_type")
		profit_currency, err := cmd.Flags().GetString("profit_currency")
		base_order_size, err := cmd.Flags().GetFloat64("base_order_size")
		safety_order_size, err := cmd.Flags().GetFloat64("safety_order_size")
		order_type, err := cmd.Flags().GetString("order_type")
		target_profit, err := cmd.Flags().GetString("target_profit")
		profit_type, err := cmd.Flags().GetString("profit_type")
		trailing_deviation, err := cmd.Flags().GetString("trailing_deviation")
		stop_loss, err := cmd.Flags().GetString("stop_loss")
		stop_loss_action, err := cmd.Flags().GetString("stop_loss_action")
		max_safety_trade_acc, err := cmd.Flags().GetString("max_safety_trade_acc")
		max_active_safety_trade_acc, err := cmd.Flags().GetString("max_active_safety_trade_acc")
		price_devation, err := cmd.Flags().GetString("price_devation")
		safety_order_volume_scale, err := cmd.Flags().GetString("safety_order_volume_scale")
		safety_order_step_scale, err := cmd.Flags().GetString("safety_order_step_scale")
		active, err := cmd.Flags().GetBool("active")
		no_deal_if_daily_volume_is_less_than, err := cmd.Flags().GetString("no_deal_if_daily_volume_is_less_than")
		min_price_to_open_deal, err := cmd.Flags().GetFloat64("min_price_to_open_deal")
		max_price_to_open_deal, err := cmd.Flags().GetFloat64("max_price_to_open_deal")
		cooldown_bewtween_deals, err := cmd.Flags().GetString("cooldown_bewtween_deals")
		open_deal_stop, err := cmd.Flags().GetString("open_deal_stop")
		user_id, err := cmd.Flags().GetString("user_id")

		if err != nil {
			return err
		}
		// Create a strategy protobuffer message
		strategy := &strategypb.Strategy{
			BotName:                       bot_name,
			SelectedStrategy:              selected_strategy,
			BotType:                       bot_type,
			Pairs:                         pairs,
			StrategyType:                  strategy_type,
			ProfitCurrency:                profit_currency,
			BaseOrderSize:                 base_order_size,
			SafetyOrderSize:               safety_order_size,
			OrderType:                     order_type,
			TargetProfit:                  target_profit,
			ProfitType:                    profit_type,
			TrailingDeviation:             trailing_deviation,
			StopLoss:                      stop_loss,
			StopLossAction:                stop_loss_action,
			MaxSafetyTradeAcc:             max_safety_trade_acc,
			MaxActiveSafetyTradeAcc:       max_active_safety_trade_acc,
			PriceDevation:                 price_devation,
			SafetyOrderVolumeScale:        safety_order_volume_scale,
			SafetyOrderStepScale:          safety_order_step_scale,
			Active:                        active,
			NoDealIfDailyVolumeIsLessThan: no_deal_if_daily_volume_is_less_than,
			MinPriceToOpenDeal:            min_price_to_open_deal,
			MaxPriceToOpenDeal:            max_price_to_open_deal,
			CooldownBewtweenDeals:         cooldown_bewtween_deals,
			OpenDealStop:                  open_deal_stop,
			UserId:                        user_id,
		}

		//record starttime of request
		startTime := time.Now()

		// RPC call
		res, err := client.CreateStrategy(
			context.Background(),
			// wrap the blog message in a CreateBlog request message
			&strategypb.CreateStrategyReq{
				Strategy: strategy,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Strategy created: %s\n", res.Strategy.Id)

		fmt.Println(res)

		diff := time.Since(startTime)
		fmt.Println("Time taken for the operation: ")
		fmt.Println(diff)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("bot_name", "a", "", "Add a bot name")
	createCmd.Flags().StringP("selected_strategy", "b", "", "Add a selected strategy")
	createCmd.Flags().StringP("bot_type", "c", "", "Add a bot type")
	createCmd.Flags().StringP("pairs", "d", "", "Add pairs")
	createCmd.Flags().StringP("strategy_type", "e", "", "Add a strategy type")
	createCmd.Flags().StringP("profit_currency", "f", "", "Add a profit currency")
	//createCmd.Flags().StringP("base_order_size", "bos", "", "Add base order size")
	createCmd.Flags().Float64P("base_order_size", "g", 0.0, "Add base order size")
	//createCmd.Flags().StringP("safety_order_size", "sos", "", "Add safety order size")
	createCmd.Flags().Float64P("safety_order_size", "i", 0.0, "Add safety order size")
	createCmd.Flags().StringP("order_type", "j", "", "Add order type")
	createCmd.Flags().StringP("target_profit", "k", "", "Add target profit")
	createCmd.Flags().StringP("profit_type", "l", "", "Add profit type")
	createCmd.Flags().StringP("trailing_deviation", "m", "", "Add trailing deviation")
	createCmd.Flags().StringP("stop_loss", "n", "", "Add stop loss")
	createCmd.Flags().StringP("stop_loss_action", "o", "", "Add stop loss action")
	createCmd.Flags().StringP("max_safety_trade_acc", "p", "", "Add max safety trade acc")
	createCmd.Flags().StringP("max_active_safety_trade_acc", "q", "", "Add max active safety trade acc")
	createCmd.Flags().StringP("price_devation", "r", "", "Add price devation")
	createCmd.Flags().StringP("safety_order_volume_scale", "s", "", "Add safety order volume scale")
	createCmd.Flags().StringP("safety_order_step_scale", "t", "", "Add safety order step scale")
	createCmd.Flags().BoolP("active", "u", true, "Add active")
	createCmd.Flags().StringP("no_deal_if_daily_volume_is_less_than", "v", "", "Add no deal if daily volume is less than")
	createCmd.Flags().Float64P("min_price_to_open_deal", "w", 0.0, "Add min price to open deal")
	createCmd.Flags().Float64P("max_price_to_open_deal", "x", 0.0, "Add max price to open deal")
	createCmd.Flags().StringP("cooldown_bewtween_deals", "y", "", "Add cooldown bewtween deals")
	createCmd.Flags().StringP("open_deal_stop", "z", "", "Add open deal stop")
	createCmd.Flags().StringP("user_id", "2", "", "Add user id")
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
