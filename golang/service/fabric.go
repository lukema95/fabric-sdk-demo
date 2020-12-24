package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maluning/fabric-sdk-sample/golang/client"
	"github.com/spf13/viper"
	"log"
	"os"
)

var(
	peerConf map[string]interface{}
	orgnizeConf map[string]interface{}
	userConf map[string]interface{}
	adminConf map[string]interface{}
	chaincodeConf map[string]interface{}
	channelConf map[string]interface{}
	configPath map[string]interface{}
)


var org1Client *client.Client

func RunFabric(){
	log.Println(" ============= 开启服务 ===============")

	peerConf = viper.GetStringMap("peer")
	orgnizeConf = viper.GetStringMap("organization")
	userConf = viper.GetStringMap("user")
	adminConf = viper.GetStringMap("admin")
	chaincodeConf = viper.GetStringMap("chaincode")
	channelConf = viper.GetStringMap("channel")
	configPath = viper.GetStringMap("path")

	org1Client = client.NewFabric()
	defer org1Client.SDK.Close()

	router := gin.New()

	resmgmt := router.Group("/resmgmt")
	{
		resmgmt.GET("/install/:version", install)
		resmgmt.GET("/instantiate/:version", instantiate)
		resmgmt.GET("/invoke", invoke)
		resmgmt.GET("/query/:key", query)
		resmgmt.GET("/delete/:key", delete)

	}

	channel := router.Group("/channel")
	{
		channel.GET("/save/:id", save)
		channel.GET("/join/:id", join)
		channel.GET("/query", queryChannel)
	}

	//msp := router.Group("/msp")
	//{
	//	msp.GET("/register/:user/:password", register)
	//	msp.GET("/enroll/:user/:password", enroll)
	//
	//}

	event := router.Group("/event")
	{
		event.GET("/register/block", registerBlockEvent)
		event.GET("/register/chaincode/:id", registerChaincodeEvent)
		event.GET("/register/tx/:id", registerTxStatusEvent)
		event.GET("/unregister/:id", unregister)
	}

	ledger := router.Group("/ledger")
	{
		ledger.GET("/query/info", queryInfo)
		ledger.GET("/query/block/:num", queryBlock)
		ledger.GET("/query/block_by_tx_id/:id", queryBlockByTxID)
		ledger.GET("/query/config", queryConfig)
		ledger.GET("/query/config_block", queryConfigBlock)
		ledger.GET("/query/transaction/:id",queryTransaction)

	}
	router.Run(":9090")
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