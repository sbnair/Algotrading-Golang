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

// listdealsbyuserCmd represents the listdealsbyuser command
var listdealsbyuserCmd = &cobra.Command{
	Use:   "listdealsbyuser",
	Short: "List all Deals by user id",
	Long:  `List all Deals in the DB by UserId`,
	RunE: func(cmd *cobra.Command, args []string) error {
		user_id, err := cmd.Flags().GetString("user_id")
		if err != nil {
			return err
		}
		// Create the request (this can be inline below too)
		req := &strategypb.ListDealReq{
			UserId: user_id,
		}

		//record starttime of request
		startTime := time.Now()

		// Call ListBlogs that returns a stream
		stream, err := client.ListDeals(context.Background(), req)
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
			//fmt.Println(res.GetDeal())
			jsonBytes, _ := json.MarshalIndent(res, "", "    ")
			fmt.Println(string(jsonBytes))
		}

		diff := time.Since(startTime)
		fmt.Println("Time taken for the operation: ")
		fmt.Println(diff)
		return nil
	},
}

func init() {
	listdealsbyuserCmd.Flags().StringP("user_id", "u", "", "The user id of the User")
	listdealsbyuserCmd.MarkFlagRequired("user_id")
	rootCmd.AddCommand(listdealsbyuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listdealsbyuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listdealsbyuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
