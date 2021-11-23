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
	"fmt"
	"os"

	"gospl/nbs"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	client      *nbs.Client
	verbose     bool
	currentData nbs.Writer
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "gospl",
	Short:         "NBS web service command line client",
	Long:          ``,
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		client = nbs.NewClient(
			viper.GetString("url"),
			viper.GetString("username"),
			viper.GetString("password"),
			viper.GetString("licence"),
		)
		client.SetVerbose(verbose)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gospl.yaml)")

	rootCmd.PersistentFlags().String("url", "https://webservices.nbs.rs", "Webservice url")
	rootCmd.PersistentFlags().String("username", "", "Authenticate with username")
	rootCmd.PersistentFlags().String("password", "", "Authenticate with password")
	rootCmd.PersistentFlags().String("licence", "", "Licence id")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Dump request and response")

	bindableFlags := []string{"url", "username", "password", "licence"}
	for _, fl := range bindableFlags {
		_ = viper.BindPFlag(fl, rootCmd.PersistentFlags().Lookup(fl))
	}

	//_ = viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	//_ = viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	//_ = viper.BindPFlag("licence", rootCmd.PersistentFlags().Lookup("licence"))
	//_ = viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gospl" (without extension).
		viper.AddConfigPath(home)
		viper.SetEnvPrefix("GOSPL_")
		viper.SetConfigName(".gospl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("# Using config file:", viper.ConfigFileUsed())
	}
}
