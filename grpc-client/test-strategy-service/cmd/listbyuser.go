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
	"io"
	"time"

	"github.com/spf13/cobra"
	strategypb "github.com/vikjdk7/Algotrading-Golang/strategy-service/proto"
)

// listbyuserCmd represents the listbyuser command
var listbyuserCmd = &cobra.Command{
	Use:   "listbyuser",
	Short: "List all Strategies by user id",
	Long:  `List all Strategies in the DB by UserId`,
	RunE: func(cmd *cobra.Command, args []string) error {
		user_id, err := cmd.Flags().GetString("user_id")
		if err != nil {
			return err
		}
		// Create the request (this can be inline below too)
		req := &strategypb.ListStrategyReq{
			UserId: user_id,
		}

		//record starttime of request
		startTime := time.Now()

		// Call ListBlogs that returns a stream
		stream, err := client.ListStrategies(context.Background(), req)
		// Check for errors
		if err != nil {
			return err
		}
		// Start iterating
		for {
			// stream.Recv returns a pointer to a ListBlogRes at the current iteration
			res, err := stream.Recv()
			// If end of stream, break the loop
			if err == io.EOF {
				break
			}
			// if err, return an error
			if err != nil {
				return err
			}
			// If everything went well use the generated getter to print the blog message
			//fmt.Println(res.GetStrategy())
			jsonBytes, _ := json.MarshalIndent(res.GetStrategy(), "", "    ")
			fmt.Println(string(jsonBytes))
		}

		diff := time.Since(startTime)
		fmt.Println("Time taken for the operation: ")
		fmt.Println(diff)
		return nil
	},
}

func init() {
	listbyuserCmd.Flags().StringP("user_id", "u", "", "The user id of the User")
	listbyuserCmd.MarkFlagRequired("user_id")
	rootCmd.AddCommand(listbyuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listbyuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listbyuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
