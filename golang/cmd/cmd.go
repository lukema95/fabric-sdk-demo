package cmd

import (
	"fmt"
	"github.com/lukema95/fabric-sdk-demo/golang/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init()  {
	serviceCmd.PersistentFlags().StringVar(&configPath, "config-path", "./config", "config file path")
	cobra.OnInitialize(InitConfig)
	viper.BindPFlag("config-path", serviceCmd.PersistentFlags().Lookup("config-path"))
	serviceCmd.AddCommand(fabricCmd)
	serviceCmd.AddCommand(fabriCACmd)
}

func InitConfig() {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


var serviceCmd = &cobra.Command{
	Use: "sdk",
	Short: "Run sdk",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify the service ")
	},

}

var fabricCmd = &cobra.Command{
	Use: "fabric",
	Short: "Start fabric service",
	Run: func(cmd *cobra.Command, args []string) {
		service.RunFabric()
	},
}

var fabriCACmd = &cobra.Command{
	Use: "ca",
	Short: "Start fabric ca service",
	Run: func(cmd *cobra.Command, args []string) {
		service.RunFabCA()
	},
}

func Execute()  {
	if err := serviceCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}