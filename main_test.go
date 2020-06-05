package main

import (
	"github.com/hyperledger/fabric-sdk-go-sample-gm/client"
	"github.com/hyperledger/fabric-sdk-go-sample-gm/cmd"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"log"
	"testing"
)

// Install, Upgrade, Instantiate, Invoke and Query ChainCode
func TestChainCode(t *testing.T) {
	log.Println("=================== Test ChainCode Begin ===================")
	defer log.Println("=================== Test ChainCode End ===================")

	cmd.InitConfig()

	peer0Org1 := "peer0.org1.example.com"

	org1Client := client.New()

	if err := org1Client.InstallCC("v1", peer0Org1); err != nil {
		//log.Panicf("Intall chaincode error: %v", err)
		t.Errorf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org1's peers")

	//InstantiateCC chaincode only need once for each channel
	if _, err := org1Client.InstantiateCC("v1", peer0Org1); err != nil {
		//log.Panicf("Instantiated chaincode error: %v", err)
		t.Errorf("Instantiated chaincode error: %v", err)
	}
	log.Println("Chaincode has been instantiated")

	if _, err := org1Client.InvokeCC([]string{peer0Org1}); err != nil {
		//log.Panicf("Invoke chaincode error: %v", err)
		t.Errorf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	if _,err := org1Client.QueryCC("peer0.org1.example.com", "a"); err != nil {
		//log.Panicf("Query chaincode error: %v", err)
		t.Errorf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org1")

	v := "v2"

	// Install new version chaincode
	if err := org1Client.InstallCC(v, peer0Org1); err != nil {
		//log.Panicf("Intall chaincode error: %v", err)
		t.Errorf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org1's peers")

	// Upgrade chaincode
	if err := org1Client.UpgradeCC(v, peer0Org1); err != nil {
		//log.Panicf("Upgrade chaincode error: %v", err)
		t.Errorf("Upgrade chaincode error: %v", err)
	}
	log.Println("Upgrade chaincode success for channel")

	if _, err := org1Client.InvokeCC([]string{"peer0.org1.example.com"}); err != nil {
		//log.Panicf("Invoke chaincode error: %v", err)
		t.Errorf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	if _, err := org1Client.QueryCC("peer0.org1.example.com", "b"); err != nil {
		//log.Panicf("Query chaincode error: %v", err)
		t.Errorf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org1")
}

// Register and Enroll User
func TestMSP(t *testing.T){
	log.Println("=================== Test MSP Begin ===================")
	defer log.Println("=================== Test MSP End ===================")

	cmd.InitConfig()

	org1Client := client.New()

	_, err := org1Client.EnrollUser("admin","adminpw")
	if err != nil {
		log.Panicf("Enroll admin failed: %s", err)
	}

	secret, err := org1Client.RegisterUser("user123", "123", "com")
	if err != nil {
		log.Panicf("Register user failed: %s", err)
	}

	_, err = org1Client.EnrollUser("user", secret)
	if err != nil {
		log.Panicf("Enroll user failed: %s", err)
	}
}

// Create, Update and Join Channel
func TestChannel(t *testing.T){
	log.Println("=================== Test Channel Begin ===================")
	defer log.Println("=================== Test Channel End ===================")

	cmd.InitConfig()

	org1Client := client.New()

	//_, err := org1Client.QueryChannels(resmgmt.WithTargets(nil))
	//if err != nil{
	//	//log.Panicf("Query channel failed :%s", err)
	//	t.Errorf("Query channel failed :%s", err)
	//}

	_, err := org1Client.SaveChannel("businesschannel", "./config/businesschannel.tx")
	if err != nil {
		//log.Panicf("Save channel failed :%s", err)
		t.Errorf("Save channel failed :%s", err)
	}

	err = org1Client.JoinChannel("businesschannel")
	if err != nil {
		//log.Panicf("Join channel failed :%s", err)
		t.Errorf("Join channel failed :%s", err)
	}

}

// Register/Unregister Block, ChainCode and txStatus
func TestEvent(t *testing.T){
	log.Println("=================== Test Event Begin ===================")
	defer log.Println("=================== Test Event End ===================")

	cmd.InitConfig()

	org1Client := client.New()

	registration, err := org1Client.RegisterBlockEvent()
	if err != nil{
		//log.Panicf("Register block even failed: %s", err)
		t.Errorf("Register block even failed: %s", err)
	}

	var ccID = "c98f65462e218ce46875ea381d8a854adfb41ad3b21a74faeec706b4ba3ea426"
	registration, err = org1Client.RegisterChaincodeEvent(ccID,"event123")
	if err != nil{
		//log.Panicf("Register chaincode event failed: %s", err)
		t.Errorf("Register chaincode event failed: %s", err)
	}

	var txID = "c98f65462e218ce46875ea381d8a854adfb41ad3b21a74faeec706b4ba3ea426"
	registration, err = org1Client.RegisterTxStatusEvent(txID)
	if err != nil{
		//log.Panicf("Regisger txStatus event failed: %s", err)
		t.Errorf("Regisger txStatus event failed: %s", err)
	}

	org1Client.UnRegister(registration)

}

// Query Block, config, Transaction and Info
func TestLedger(t *testing.T){
	log.Println("=================== Test Ledger Begin ===================")
	defer log.Println("=================== Test Ledger End ===================")

	cmd.InitConfig()

	org1Client := client.New()

	_, err := org1Client.QueryInfo()
	if err != nil{
		//log.Panicf("QueryInfo failed: %s", err)
		t.Errorf("QueryInfo failed: %s", err)
	}

	_, err = org1Client.QueryBlock(1)
	if err != nil{
		//log.Panicf("QueryBlock failed: %s", err)
		t.Errorf("QueryBlock failed: %s", err)
	}

	var txID = "c98f65462e218ce46875ea381d8a854adfb41ad3b21a74faeec706b4ba3ea426"
	_, err = org1Client.QueryTransaction(fab.TransactionID(txID))
	if err != nil{
		//log.Panicf("QueryTransaction failed: %s", err)
		t.Errorf("QueryTransaction failed: %s", err)
	}

	_, err = org1Client.QueryConfigBlock()
	if err != nil{
		//log.Panicf("QueryConfigBlock failed: %s", err)
		t.Errorf("QueryConfigBlock failed: %s", err)
	}


	_, err = org1Client.QueryConfig()
	if err != nil{
		//log.Panicf("QueryConfig failed: %s", err)
		t.Errorf("QueryConfig failed: %s", err)
	}

}
