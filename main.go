package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	env := flag.String("env", "production", "please set the environment variables properly")
	flag.Parse()

	if *env == "development" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	UseController(router)

	fmt.Printf("run server successful. please access http://127.0.0.1:%s\n", GlobalConfig.PORT)

	if err := router.Run(fmt.Sprintf(":%s", GlobalConfig.PORT)); err != nil {
		fmt.Printf("failed to start server:%v\n", err)
		return
	}
}
