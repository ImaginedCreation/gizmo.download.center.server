package main

import (
	"github.com/gin-gonic/gin"
)

func __Register(c *gin.Context) {
	var b Register_P
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(int(RESCODEERROR), RES_ERROR_FN(ParseValidatorError(err)))
		return
	}
	cx := c.Request.Context()
	if err := Register(&b, &cx); err != nil {
		c.JSON(int(RESCODEERROR), RES_ERROR_FN(err.Error()))
		return
	}
	c.JSON(int(RESCODEOK), RES_OK_FN("login was successful"))
}

func __Login(c *gin.Context) {
	var b Login_P
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(int(RESCODEERROR), RES_ERROR_FN(ParseValidatorError(err)))
		return
	}
	cx := c.Request.Context()
	var user USER
	if err := Login(&b, &user, &cx); err != nil {
		c.JSON(int(RESCODEERROR), RES_ERROR_FN(err.Error()))
		return
	}
	token, err := GenerateToken(&Token{USERNAME: user.UserName, VERSION: user.Version})
	if err != nil {
		c.JSON(int(RESCODEERROR), RES_ERROR_FN("login failed"))
		return
	}
	c.JSON(int(RESCODEOK), RES_OK_FN[string](token))
}

func UseController(router *gin.Engine) {
	user_router := router.Group("user")
	{
		user_router.POST("register", __Register)
		user_router.POST("login", __Login)
	}
}
