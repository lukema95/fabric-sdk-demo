package service

import (
	"github.com/gin-gonic/gin"
)

func install(ctx *gin.Context){
	version, _ := ctx.Params.Get("version")
	err := org1Client.InstallCC(version, peerConf["org1"].(string))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	}else {
		ctx.JSON(200, gin.H{
			"result": "install chaincode successful",
		})
	}

}

func instantiate(ctx *gin.Context){
	version, _ := ctx.Params.Get("version")
	txID, err := org1Client.InstantiateCC(version,peerConf["org1"].(string))
	if err != nil{
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	}else {
		ctx.JSON(200, gin.H{
			"result": "initiate chaincode successful",
			"txID" : txID,
		})
	}
}

func invoke(ctx *gin.Context){
	resp, err := org1Client.InvokeCC([]string{peerConf["org1"].(string)})
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else {
		ctx.JSON(200, gin.H{
			"validate": resp.TxValidationCode,
			"txID": resp.TransactionID,
			"status": resp.ChaincodeStatus,
		})
	}
}

func delete(ctx *gin.Context){
	key, _ := ctx.Params.Get("key")
	resp, err := org1Client.InvokeCCDelete([]string{peerConf["org1"].(string)},key)
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"validate": resp.TxValidationCode,
			"txID": resp.TransactionID,
			"status": resp.ChaincodeStatus,
		})
	}
}

func query(ctx *gin.Context){
	key, _ := ctx.Params.Get("key")
	resp, err := org1Client.QueryCC(peerConf["org1"].(string),key)
	if err != nil{
		 ctx.JSON(500, gin.H{
		 	"error" : resp.Responses,
		})
	}else {
		ctx.JSON(200, gin.H{
			"txID":   resp.TransactionID,
			"result": string(resp.Payload),
		})
	}
}

