package main

import (
	"docker-example/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	controller.Init(router)
	router.Run(":8000")

}
