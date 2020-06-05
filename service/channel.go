package service

import "github.com/gin-gonic/gin"

func save(ctx *gin.Context){
	id, _ := ctx.Params.Get("id")
	resp, err := org1Client.SaveChannel(id, (channelConf["config_path"]).(string))
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else {
		ctx.JSON(200, gin.H{
			"result": "save channel successful",
			"txID": resp,
		})
	}
}

func join(ctx *gin.Context){
	id, _ := ctx.Params.Get("id")
	err := org1Client.JoinChannel(id)
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else{
		ctx.JSON(200, gin.H{
			"result": "join channel successful",
		})
	}
}

func queryChannel(ctx *gin.Context){
	resp, err := org1Client.QueryChannels()
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