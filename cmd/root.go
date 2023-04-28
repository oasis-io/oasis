package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		viper.AddConfigPath("./")
		viper.SetConfigType("yaml")
		viper.SetConfigName("oasis")
	}

	//viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
