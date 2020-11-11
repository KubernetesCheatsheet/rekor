/*
Copyright © 2020 Luke Hinds <lhinds@redhat.com>

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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "rekor",
	Short: "Rekor CLI",
	Long:  `Rekor command line interface tool`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rekor.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().String("rekor_server", "http://localhost:3000", "Server address:port")
	viper.BindPFlag("rekor_server", rootCmd.PersistentFlags().Lookup("rekor_server"))

	rootCmd.PersistentFlags().String("rekord", "", "Rekor rekord file")
	viper.BindPFlag("rekord", rootCmd.PersistentFlags().Lookup("rekord"))

	rootCmd.PersistentFlags().String("signature", "", "Rekor signature")
	viper.BindPFlag("signature", rootCmd.PersistentFlags().Lookup("signature"))

	rootCmd.PersistentFlags().String("public_key", "", "Rekor publickey")
	viper.BindPFlag("public_key", rootCmd.PersistentFlags().Lookup("public_key"))

	rootCmd.PersistentFlags().String("artifact_path", "", "Rekor artifact path")
	viper.BindPFlag("artifact_path", rootCmd.PersistentFlags().Lookup("artifact_path"))

	rootCmd.PersistentFlags().String("artifact_url", "", "Rekor artifact url")
	viper.BindPFlag("artifact_url", rootCmd.PersistentFlags().Lookup("artifact_url"))

	rootCmd.PersistentFlags().String("artifact_sha", "", "Rekor artifact sha")
	viper.BindPFlag("artifact_sha", rootCmd.PersistentFlags().Lookup("artifact_sha"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".rekor")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
