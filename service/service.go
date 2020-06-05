package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go-sample-gm/client"
	"log"
)

var (
	peer0Org1 = "peer0.org1.example.com"
	peer0Org2 = "peer0.org2.example.com"

	userName = "user12"
	password = "pw"
	department = "com"

	channelID = "businesschannel"
	channelConfigPath= "./config/businesschannel.tx"

	eventFilter = "event123"
	txID = "21bc37e36c753bc67b098e808e0993d4c11815e51dfc2b5f2eb50dab28a91c85"
)

var org1Client *client.Client
var org2Client *client.Client

func Run(){
	log.Println(" ============= 开启服务 ===============")
	org1Client = client.New("./config/org1-config.yaml", "Org1", "Admin", "User1")
	org2Client = client.New("./config/org2-config.yaml", "Org2", "Admin", "User1")

	defer org1Client.SDK.Close()
	defer org2Client.SDK.Close()

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

	msp := router.Group("/msp")
	{
		msp.GET("/register/:user/:password", register)
		msp.GET("/enroll/:user/:password", enroll)

	}

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

