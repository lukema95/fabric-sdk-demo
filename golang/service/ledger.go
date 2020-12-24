package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"log"
	"strconv"
)

func queryInfo(ctx *gin.Context){
	resp, err := org1Client.QueryInfo()
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": resp,
		})
	}
}

func queryBlock(ctx *gin.Context){
	num, _ := ctx.Params.Get("num")
	resp, err := org1Client.QueryBlock(stringToUint64(num))
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": resp,
		})
	}
}

func queryBlockByTxID(ctx *gin.Context){
	id, _ := ctx.Params.Get("id")
	resp, err := org1Client.QueryBlockByTxID(fab.TransactionID(id))
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": resp,
		})
	}
}

func queryConfig(ctx *gin.Context){
	resp, err := org1Client.QueryConfig()
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": resp,
		})
	}
}

func queryConfigBlock(ctx *gin.Context){
	resp, err := org1Client.QueryConfigBlock()
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": resp,
		})
	}
}

func queryTransaction(ctx *gin.Context){
	id, _ := ctx.Params.Get("id")
	resp, err := org1Client.QueryTransaction(fab.TransactionID(id))
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": resp,
		})
	}
}

func stringToUint64(str string) uint64 {
	num, err := strconv.Atoi(str)
	if err != nil{
		log.Panicf("can not transform %d to uint64", str)
	}
	return uint64(num)
}