package service

import "github.com/gin-gonic/gin"

func register(ctx *gin.Context){
	user, _ := ctx.Params.Get("user")
	pwd, _ := ctx.Params.Get("password")

	secret, err := org1Client.RegisterUser(user,pwd,department)
	if err != nil{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else {
		ctx.JSON(200, gin.H{
			"result": "register user successful",
			"secret": secret,
		})
	}
}

func enroll(ctx *gin.Context){
	user, _ := ctx.Params.Get("user")
	pwd, _ := ctx.Params.Get("password")

	ok, err := org1Client.EnrollUser(user, pwd)
	if !ok{
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}else {
		ctx.JSON(200, gin.H{
			"result": "enroll user successful",
		})
	}
}