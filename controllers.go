package main

import "github.com/gin-gonic/gin"

func __Register(c *gin.Context) {
	var b TRegister_P
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(int(RESCODEERROR), RES_ERROR_FN(ParseValidatorError(err)))
		return
	}
	c.JSON(int(RESCODEOK), RES_OK_FN("ok"))
}

func __Login(c *gin.Context) {
	var b TLogin_P
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(int(RESCODEERROR), RES_ERROR_FN(ParseValidatorError(err)))
		return
	}
	token := "token-123"
	c.JSON(int(RESCODEOK), RES_OK_FN[string](token))
}

func UseController(router *gin.Engine) {
	user_router := router.Group("928584ec-0014-5d05-7dbd-e51a27b4d358")
	{
		user_router.POST("register", __Register)
		user_router.POST("login", __Login)
	}
}
