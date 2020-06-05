package main

import (
	"github.com/hyperledger/fabric-sdk-go-sample-gm/client"
	"github.com/hyperledger/fabric-sdk-go-sample-gm/cmd"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
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
func main() {

	cmd.Execute()
	//org1Client := client.New("./config/org1-config.yaml", "Org1", "Admin", "User1")
	//org2Client := client.New("./config/org2-config.yaml", "Org2", "Admin", "User1")
	//
	//defer org1Client.SDK.Close()
	//defer org2Client.SDK.Close()
	//
	//// Install, Upgrade, Instantiate, Invoke and Query ChainCode
	//TestChainCode(org1Client, org2Client)
	//
	//// Register and Enroll User
	//TestMSP(org1Client, peer0Org1, password, department)
	//
	//// Create, Update and Join Channel
	//TestChannel(org1Client,channelID,channelConfigPath)
	//
	//// Register/Unregister Block, ChainCode and txStatus
	//TestEvent(org1Client, eventFilter, org1Client.ChannelID, txID)
	//
	//// Query Block, config, Transaction and Info
	//TestLedger(org1Client, fab.TransactionID(txID), 1)
}

// Install, Upgrade, Instantiate, Invoke and Query ChainCode
func TestChainCode(cli1, cli2 *client.Client) {
	log.Println("=================== Test ChainCode Begin ===================")
	defer log.Println("=================== Test ChainCode End ===================")

	if err := cli1.InstallCC("v1", peer0Org1); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org1's peers")

	if err := cli2.InstallCC("v1", peer0Org2); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org2's peers")

	//InstantiateCC chaincode only need once for each channel
	if _, err := cli1.InstantiateCC("v1", peer0Org1); err != nil {
		log.Panicf("Instantiated chaincode error: %v", err)
	}
	log.Println("Chaincode has been instantiated")

	if _, err := cli1.InvokeCC([]string{peer0Org1}); err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	if _,err := cli1.QueryCC("peer0.org1.example.com", "a"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org1")

	v := "v2"

	// Install new version chaincode
	if err := cli1.InstallCC(v, peer0Org1); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org1's peers")

	if err := cli2.InstallCC(v, peer0Org2); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org2's peers")

	// Upgrade chaincode only need once for each channel
	if err := cli1.UpgradeCC(v, peer0Org2); err != nil {
		log.Panicf("Upgrade chaincode error: %v", err)
	}
	log.Println("Upgrade chaincode success for channel")

	if _, err := cli1.InvokeCC([]string{"peer0.org1.example.com",
		"peer0.org2.example.com"}); err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	if _, err := cli1.QueryCC("peer0.org2.example.com", "a"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org2")
}

// Register and Enroll User
func TestMSP(client *client.Client, name, password, department string){
	log.Println("=================== Test MSP Begin ===================")
	defer log.Println("=================== Test MSP End ===================")

	_, err := client.EnrollUser("admin","adminpw")
	if err != nil {
		log.Panicf("Enroll admin failed: %s", err)
	}

	secret, err := client.RegisterUser(name, password, department)
	if err != nil {
		log.Panicf("Register user failed: %s", err)
	}

	_, err = client.EnrollUser(name, secret)
	if err != nil {
		log.Panicf("Enroll user failed: %s", err)
	}
}

// Create, Update and Join Channel
func TestChannel(client *client.Client, channelId string, channelConfigPath string){
	log.Println("=================== Test Channel Begin ===================")
	defer log.Println("=================== Test Channel End ===================")

	_, err := client.QueryChannels(resmgmt.WithTargets(nil))
	if err != nil{
		log.Panicf("Query channel failed :%s", err)
	}

	_, err = client.SaveChannel(channelId, channelConfigPath)
	if err != nil {
		log.Panicf("Save channel failed :%s", err)
	}

	err = client.JoinChannel(channelId)
	if err != nil {
		log.Panicf("Join channel failed :%s", err)
	}

}

// Register/Unregister Block, ChainCode and txStatus
func TestEvent(client *client.Client, eventFilter, ccID, txID string){
	log.Println("=================== Test Event Begin ===================")
	defer log.Println("=================== Test Event End ===================")
	registration, err := client.RegisterBlockEvent()
	if err != nil{
		log.Panicf("Register block even failed: %s", err)
	}

	registration, err = client.RegisterChaincodeEvent(ccID,eventFilter)
	if err != nil{
		log.Panicf("Register chaincode event failed: %s", err)
	}

	registration, err = client.RegisterTxStatusEvent(txID)
	if err != nil{
		log.Panicf("Regisger txStatus event failed: %s", err)
	}

	client.UnRegister(registration)

}

// Query Block, config, Transaction and Info
func TestLedger(client *client.Client,txID fab.TransactionID, blockNum uint64, options ...ledger.RequestOption ){
	log.Println("=================== Test Ledger Begin ===================")
	defer log.Println("=================== Test Ledger End ===================")

	_, err := client.QueryInfo(options...)
	if err != nil{
		log.Panicf("QueryInfo failed: %s", err)
	}

	_, err = client.QueryBlock(blockNum, options...)
	if err != nil{
		log.Panicf("QueryBlock failed: %s", err)
	}

	_, err = client.QueryTransaction(txID, options...)
	if err != nil{
		log.Panicf("QueryTransaction failed: %s", err)
	}

	_, err = client.QueryConfigBlock(options...)
	if err != nil{
		log.Panicf("QueryConfigBlock failed: %s", err)
	}


	_, err = client.QueryConfig(options...)
	if err != nil{
		log.Panicf("QueryConfig failed: %s", err)
	}

}