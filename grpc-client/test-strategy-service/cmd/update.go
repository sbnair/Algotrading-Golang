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
	strategypb "github.com/vikjdk7/Algotrading-Golang/strategy-service/proto"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a strategy by its Id",
	Long:  `Update a Strategy by its Id in MongoDB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the flags from CLI
		id, err := cmd.Flags().GetString("id")
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

		// Create an UpdateExchangeReq
		strategy := &strategypb.Strategy{
			Id:                            id,
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
		}

		res, err := client.UpdateStrategy(context.Background(),
			&strategypb.UpdateStrategyReq{
				Strategy: strategy,
			})
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("id", "1", "", "The id of the exchange")
	updateCmd.Flags().StringP("bot_name", "a", "", "Add a bot name")
	updateCmd.Flags().StringP("selected_strategy", "b", "", "Add a selected strategy")
	updateCmd.Flags().StringP("bot_type", "c", "", "Add a bot type")
	updateCmd.Flags().StringP("pairs", "d", "", "Add pairs")
	updateCmd.Flags().StringP("strategy_type", "e", "", "Add a strategy type")
	updateCmd.Flags().StringP("profit_currency", "f", "", "Add a profit currency")
	//updateCmd.Flags().StringP("base_order_size", "bos", "", "Add base order size")
	updateCmd.Flags().Float64P("base_order_size", "g", 0.0, "Add base order size")
	//updateCmd.Flags().StringP("safety_order_size", "sos", "", "Add safety order size")
	updateCmd.Flags().Float64P("safety_order_size", "i", 0.0, "Add safety order size")
	updateCmd.Flags().StringP("order_type", "j", "", "Add order type")
	updateCmd.Flags().StringP("target_profit", "k", "", "Add target profit")
	updateCmd.Flags().StringP("profit_type", "l", "", "Add profit type")
	updateCmd.Flags().StringP("trailing_deviation", "m", "", "Add trailing deviation")
	updateCmd.Flags().StringP("stop_loss", "n", "", "Add stop loss")
	updateCmd.Flags().StringP("stop_loss_action", "o", "", "Add stop loss action")
	updateCmd.Flags().StringP("max_safety_trade_acc", "p", "", "Add max safety trade acc")
	updateCmd.Flags().StringP("max_active_safety_trade_acc", "q", "", "Add max active safety trade acc")
	updateCmd.Flags().StringP("price_devation", "r", "", "Add price devation")
	updateCmd.Flags().StringP("safety_order_volume_scale", "s", "", "Add safety order volume scale")
	updateCmd.Flags().StringP("safety_order_step_scale", "t", "", "Add safety order step scale")
	updateCmd.Flags().BoolP("active", "u", true, "Add active")
	updateCmd.Flags().StringP("no_deal_if_daily_volume_is_less_than", "v", "", "Add no deal if daily volume is less than")
	updateCmd.Flags().Float64P("min_price_to_open_deal", "w", 0.0, "Add min price to open deal")
	updateCmd.Flags().Float64P("max_price_to_open_deal", "x", 0.0, "Add max price to open deal")
	updateCmd.Flags().StringP("cooldown_bewtween_deals", "y", "", "Add cooldown bewtween deals")
	updateCmd.Flags().StringP("open_deal_stop", "z", "", "Add open deal stop")
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
