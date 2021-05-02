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
	pricepb "github.com/vikjdk7/Algotrading-Golang/price-service/proto"
)

// listassetbysymbolCmd represents the listassetbysymbol command
var listassetbysymbolCmd = &cobra.Command{
	Use:   "listassetbysymbol",
	Short: "List Asset By Symbol",
	Long:  `List Asset by Symbol`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, err := cmd.Flags().GetString("symbol")
		if err != nil {
			return err
		}

		req := &pricepb.ListAssetBySymbolReq{
			Symbol: symbol,
		}

		//record starttime of request
		startTime := time.Now()

		stream, err := client.ListAssetBySymbol(context.Background(), req)
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
		diff := time.Since(startTime)
		fmt.Print("Time taken for the operation: ")
		fmt.Println(diff)
		return nil
	},
}

func init() {
	listassetbysymbolCmd.Flags().StringP("symbol", "s", "", "The symbol of the Asset")
	listassetbysymbolCmd.MarkFlagRequired("symbol")
	rootCmd.AddCommand(listassetbysymbolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listbysymbolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listbysymbolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
