package client

import (
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type Client struct {
	// Fabric network information
	ConfigPath string
	OrgName    string
	OrgAdmin   string
	OrgUser    string

	// Sdk clients
	SDK 			*fabsdk.FabricSDK
	resMgtClient  	*resmgmt.Client
	channelClient  	*channel.Client
	mspClient 		*msp.Client
	ledgerClient 	*ledger.Client
	eventClient 	*event.Client

	// Same for each peer
	ChannelID string
	CCID      string // chaincode ID
	CCPath    string // chaincode source path
	CCGoPath  string // GOPATH used for chaincode
}

func New() *Client {
	orgnizeConf := viper.GetStringMap("organization")
	userConf := viper.GetStringMap("user")
	adminConf := viper.GetStringMap("admin")
	chaincodeConf := viper.GetStringMap("chaincode")
	channelConf := viper.GetStringMap("channel")
	configPath := viper.GetStringMap("path")

	c := &Client{
		ConfigPath: configPath["org1"].(string),
		OrgName:    orgnizeConf["org1"].(string),
		OrgAdmin:   adminConf["name"].(string),
		OrgUser:    userConf["name"].(string),

		CCID:      chaincodeConf["id"].(string),
		CCPath:    chaincodeConf["path"].(string),
		CCGoPath:  os.Getenv("GOPATH"),
		ChannelID: channelConf["id"].(string),
	}

	// create sdk
	sdk, err := fabsdk.New(config.FromFile(c.ConfigPath))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}
	c.SDK = sdk
	log.Println("Initialized fabric sdk")

	// create resmgmt client
	rctx := sdk.Context(fabsdk.WithUser(c.OrgAdmin), fabsdk.WithOrg(c.OrgName))
	c.resMgtClient, err = resmgmt.New(rctx)
	if err != nil {
		log.Panicf("failed to create resource client: %s", err)
	}
	log.Println("Initialized resource client")

	// create channel client
	cctx := sdk.ChannelContext(c.ChannelID, fabsdk.WithUser(c.OrgUser))
	c.channelClient, err = channel.New(cctx)
	if err != nil {
		log.Panicf("failed to create channel client: %s", err)
	}
	log.Println("Initialized channel client")

	// create msp client
	mctx := sdk.Context(fabsdk.WithUser(c.OrgAdmin), fabsdk.WithOrg(c.OrgName))
	c.mspClient, err = msp.New(mctx)
	if err != nil{
		log.Panicf("failed to create msp client: %s", err)
	}
	log.Println("Initialized msp client")

	// create event client
	ectx := sdk.ChannelContext(c.ChannelID, fabsdk.WithUser(c.OrgUser))
	c.eventClient, err = event.New(ectx,event.WithBlockEvents())
	if err != nil{
		log.Panicf("failed to create even client: %s", err)
	}
	log.Println("Initialized even client")

	// create ledger client
	lctx := sdk.ChannelContext(c.ChannelID, fabsdk.WithUser(c.OrgAdmin))
	c.ledgerClient, err = ledger.New(lctx)
	if err != nil{
		log.Panicf("failed to create ledger client: %s", err)
	}
	log.Println("Initialized ledger client")

	return c
}
