package cmd

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go-sample-gm/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init()  {
	serviceCmd.PersistentFlags().StringVar(&configPath, "config-path", "./config", "config file path")
	cobra.OnInitialize(InitConfig)
	viper.BindPFlag("config-path", serviceCmd.PersistentFlags().Lookup("config-path"))
	serviceCmd.AddCommand(startCmd)
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
	Use: "service",
	Short: "Run service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify the service ")
	},

}

var startCmd = &cobra.Command{
	Use: "start",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		service.Run()
	},
}

func Execute()  {
	if err := serviceCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}