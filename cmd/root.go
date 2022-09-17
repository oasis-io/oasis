package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	yaml3 "gopkg.in/yaml.v3"
	"io/ioutil"
	"oasis/config"
	"os"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "oasis",
	Short: "Oasis",
	Long:  `Oasis Database DevOps v0.1.0`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "config file path")
}

func initConfig() {
	var cfg = &config.Config{}

	if configPath != "" {
		b, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Errorf("load config path: %s failed error :%v", configPath, err)
		}

		err = yaml3.Unmarshal(b, cfg)
		if err != nil {
			fmt.Errorf("unmarshal config file error %v", err)
		}
	}

	Run(cfg)
}
