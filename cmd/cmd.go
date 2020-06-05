package cmd

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go-sample-gm/service"
	"github.com/spf13/cobra"
	"os"
)

var configPath string

func init()  {
	serviceCmd.AddCommand(startCmd)
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