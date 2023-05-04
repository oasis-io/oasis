package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"oasis/pkg/log"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "oasis",
		Short: "Oasis 数据库运维平台",
		Long:  `Oasis是一个数据库运维平台，可以基于K8S 创建、维护、管理MySQL、Redis等数据库.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./oasis.yaml)")
	viper.SetDefault("license", "Apache License 2.0")

}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		//home, err := os.UserHomeDir()
		//cobra.CheckErr(err)
		//viper.AddConfigPath(home)

		viper.AddConfigPath(".")
		viper.SetConfigType("toml")
		viper.SetConfigName("oasis")

		// server default value
		viper.SetDefault("server.bind", "127.0.0.1")
		viper.SetDefault("server.port", "9590")
		viper.SetDefault("server.error_log", "./oasis.log")
		viper.SetDefault("server.access_log", "./access.log")

	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		log.Info("Initializing config file")
		log.Info("Using config file", zap.String("file", viper.ConfigFileUsed()))
	}
}
