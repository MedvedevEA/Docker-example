package main

import (
	"docker/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	controller.Init(router)
	router.Run(":8000")

}
