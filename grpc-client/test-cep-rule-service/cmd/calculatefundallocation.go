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
	ceprulepb "github.com/vikjdk7/Algotrading-Golang/cep-rule-service/proto"
)

// calculatefundallocationCmd represents the calculatefundallocation command
var calculatefundallocationCmd = &cobra.Command{
	Use:   "calculatefundallocation",
	Short: "Calculate Fund allocation for Strategy",
	Long:  `Calculate Fund allocation for Strategy.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		base_order_size, err := cmd.Flags().GetFloat64("base_order_size")
		safety_order_size, err := cmd.Flags().GetFloat64("safety_order_size")
		safety_order_volume_scale, err := cmd.Flags().GetFloat64("safety_order_volume_scale")
		safety_order_step_scale, err := cmd.Flags().GetFloat64("safety_order_step_scale")
		max_active_safety_trade_count, err := cmd.Flags().GetFloat64("max_active_safety_trade_count")
		total_no_deals, err := cmd.Flags().GetInt64("total_no_deals")

		if err != nil {
			return err
		}

		calculateStrategyFundAllocationReq := &ceprulepb.CalculateStrategyFundAllocationReq{
			BaseOrderSize:             base_order_size,
			SafetyOrderSize:           safety_order_size,
			SafetyOrderVolumeScale:    safety_order_volume_scale,
			SafetyOrderStepScale:      safety_order_step_scale,
			MaxActiveSafetyTradeCount: max_active_safety_trade_count,
			TotalNoDeals:              total_no_deals,
		}
		//record starttime of request
		startTime := time.Now()

		res, err := client.CalculateStrategyFundAllocation(
			context.Background(),
			// wrap the blog message in a CreateBlog request message
			calculateStrategyFundAllocationReq,
		)
		if err != nil {
			return err
		}
		//fmt.Printf("Total Funds for Allocation: %v\n", res.TotalFundForAllocation)
		jsonBytes, _ := json.MarshalIndent(res, "", "    ")
		fmt.Println(string(jsonBytes))
		diff := time.Since(startTime)
		fmt.Print("Time taken for the operation: ")
		fmt.Println(diff)

		return nil
	},
}

func init() {
	calculatefundallocationCmd.Flags().Float64P("base_order_size", "b", 0.0, "Add base order size")
	calculatefundallocationCmd.Flags().Float64P("safety_order_size", "s", 0.0, "Add safety order size")
	calculatefundallocationCmd.Flags().Float64P("safety_order_volume_scale", "v", 0.0, "Add base order size")
	calculatefundallocationCmd.Flags().Float64P("safety_order_step_scale", "o", 0.0, "Add safety order size")
	calculatefundallocationCmd.Flags().Float64P("max_active_safety_trade_count", "a", 0.0, "Add base order size")
	calculatefundallocationCmd.Flags().Int64P("total_no_deals", "d", 0, "Add safety order size")
	calculatefundallocationCmd.MarkFlagRequired("base_order_size")
	calculatefundallocationCmd.MarkFlagRequired("safety_order_size")
	calculatefundallocationCmd.MarkFlagRequired("safety_order_volume_scale")
	calculatefundallocationCmd.MarkFlagRequired("total_no_deals")
	calculatefundallocationCmd.MarkFlagRequired("max_active_safety_trade_count")
	rootCmd.AddCommand(calculatefundallocationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculatefundallocationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calculatefundallocationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
