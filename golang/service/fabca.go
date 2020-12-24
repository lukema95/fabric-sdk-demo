package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"../client"
	"github.com/spf13/viper"
	"log"
	"os"
)

func RunFabCA(){
	log.Println(" ============= 开启服务 ===============")

	peerConf = viper.GetStringMap("peer")
	orgnizeConf = viper.GetStringMap("organization")
	userConf = viper.GetStringMap("user")
	adminConf = viper.GetStringMap("admin")
	chaincodeConf = viper.GetStringMap("chaincode")
	channelConf = viper.GetStringMap("channel")
	configPath = viper.GetStringMap("path")

	org1Client = client.NewFabCA()
	defer org1Client.SDK.Close()

	router := gin.New()

	msp := router.Group("/msp")
	{
		msp.GET("/register/:user/:password", register)
		msp.GET("/enroll/:user/:password", enroll)

	}

	router.Run(":9091")
}

func init() {
	configPath := "./config"
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}