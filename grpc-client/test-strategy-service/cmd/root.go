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
	"log"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	strategypb "github.com/vikjdk7/Algotrading-Golang/strategy-service/proto"
	"google.golang.org/grpc"
)

var cfgFile string

// Client and context global vars
var client strategypb.StrategyServiceClient
var requestCtx context.Context
var requestOpts grpc.DialOption

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test-strategy-service",
	Short: "a gRPC client to communicate with the StrategyService server",
	Long: `a gRPC client to communicate with the StrategyService server.
	You can use this client to create and read Strategy.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test-strategy-service.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// After Cobra root config init, initialize the client
	fmt.Println("Starting Strategy Service Client")
	// Establish context to timeout after 10 seconds if server does not respond
	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	// Establish insecure grpc options (no TLS)
	requestOpts = grpc.WithInsecure()
	// Dial the server, returns a client connection
	//conn, err := grpc.Dial("localhost:50052", requestOpts)
	conn, err := grpc.Dial("130.211.61.141:50052", requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50052: %v", err)
	}
	// Instantiate the BlogServiceClient with our client connection to the server
	client = strategypb.NewStrategyServiceClient(conn)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".test-strategy-service" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".test-strategy-service")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
