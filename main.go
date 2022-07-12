package main

import (
	"main/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	helpers.Initial_data()
	api := gin.Default()
	api.GET("/random/mean", helpers.MeanEndpoint)
	api.Run("localhost:8080")
}
