package service

import "github.com/gin-gonic/gin"

func registerBlockEvent(ctx *gin.Context){
	_, err := org1Client.RegisterBlockEvent()
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": " Register block event successful",
		})
	}
}

func registerChaincodeEvent(ctx *gin.Context){
	ccID, _ := ctx.Params.Get("id")
	_, err := org1Client.RegisterChaincodeEvent(ccID,"event123")
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": "Register Chaincode event successful",
		})
	}
}
func registerTxStatusEvent(ctx *gin.Context){
	id, _ := ctx.Params.Get("id")
	_, err := org1Client.RegisterTxStatusEvent(id)
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": "Register txStatus event successful",
		})
	}
}

func unregister(ctx *gin.Context){
	id, _ := ctx.Params.Get("id")
	org1Client.UnRegister(id)
	ctx.JSON(200, gin.H{
		"result": "unregister successful",
	})
}