package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"oasis/app"
	"oasis/config"
	"runtime"
	"time"
)

var cfgFile string

func init() {
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./oasis.toml)")
	viper.SetDefault("license", "Apache License 2.0")

	// Add version command
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "oasis",
	Short: "Oasis database ops",
	Long:  `Oasis is a tool for managing databases.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.InitConfig(cfgFile); err != nil { // check config file
			fmt.Println("Error initializing config:", err)
			return
		}

		app.RunServer() // start web server
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Oasis",
	Long:  `All software has versions. This is Oasis's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oasis version", config.VERSION)
		fmt.Println("build date:", time.Now().Format("2006-01-02"))
		fmt.Println("go version:", runtime.Version())
		fmt.Println("GOOS:", runtime.GOOS)
		fmt.Println("GOARCH:", runtime.GOARCH)
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
