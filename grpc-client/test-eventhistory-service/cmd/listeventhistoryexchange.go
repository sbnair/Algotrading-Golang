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

	"github.com/spf13/cobra"
	eventhistorypb "github.com/vikjdk7/Algotrading-Golang/eventhistory-service/proto"
)

// listeventhistoryexchangeCmd represents the listeventhistoryexchange command
var listeventhistoryexchangeCmd = &cobra.Command{
	Use:   "listeventhistoryexchange",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		user_id, err := cmd.Flags().GetString("user_id")
		if err != nil {
			return err
		}

		req := &eventhistorypb.ListEventHistoryExchangeReq{
			UserId: user_id,
		}
		// Call ListBlogs that returns a stream
		stream, err := client.ListEventHistoryExchange(context.Background(), req)
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
			//fmt.Println(res)
			// If everything went well use the generated getter to print the exchange message
			jsonBytes, _ := json.MarshalIndent(res, "", "    ")
			fmt.Println(string(jsonBytes))
		}
		return nil
	},
}

func init() {
	listeventhistoryexchangeCmd.Flags().StringP("user_id", "u", "", "The user id of the User")
	listeventhistoryexchangeCmd.MarkFlagRequired("user_id")
	rootCmd.AddCommand(listeventhistoryexchangeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listeventhistoryexchangeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listeventhistoryexchangeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
