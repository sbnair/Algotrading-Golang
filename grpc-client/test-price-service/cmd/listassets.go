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
	pricepb "github.com/vikjdk7/Algotrading-Golang/price-service/proto"
)

// listassetsCmd represents the listassets command
var listassetsCmd = &cobra.Command{
	Use:   "listassets",
	Short: "List All Assets",
	Long:  `List All the Assets/Stocks`,
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &pricepb.ListAssetsReq{}
		stream, err := client.ListAssets(context.Background(), req)
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
			// If everything went well use the generated getter to print the exchange message
			//fmt.Println(res.GetPosition())
			jsonBytes, _ := json.MarshalIndent(res.GetAsset(), "", "    ")
			fmt.Println(string(jsonBytes))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listassetsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listassetsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listassetsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
