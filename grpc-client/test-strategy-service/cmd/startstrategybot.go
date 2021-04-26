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

	"github.com/spf13/cobra"
	strategypb "github.com/vikjdk7/Algotrading-Golang/strategy-service/proto"
)

// startstrategybotCmd represents the startstrategybot command
var startstrategybotCmd = &cobra.Command{
	Use:   "startstrategybot",
	Short: "Start the strategy bot",
	Long:  `Start a strategybot for the selected stocks and run their deals.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		strategy_id, err := cmd.Flags().GetString("strategy_id")
		stock, err := cmd.Flags().GetStringSlice("stock")
		if err != nil {
			return err
		}

		deal := make([]*strategypb.Stock, len(stock))
		for i, v := range stock {
			deal[i] = &strategypb.Stock{
				StockName: v,
			}
		}

		res, err := client.StartBot(
			context.Background(),
			// wrap the blog message in a CreateBlog request message
			&strategypb.StartBotReq{
				StrategyId: strategy_id,
				Stocks:     deal,
			},
		)

		if err != nil {
			return err
		}

		//fmt.Println(res)
		jsonBytes, _ := json.MarshalIndent(res, "", "    ")
		fmt.Println(string(jsonBytes))

		return nil
	},
}

func init() {
	startstrategybotCmd.Flags().StringP("strategy_id", "i", "", "Add a strategy name")
	startstrategybotCmd.Flags().StringSliceP("stock", "d", []string{}, "Add stock")
	rootCmd.AddCommand(startstrategybotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startstrategybotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startstrategybotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
