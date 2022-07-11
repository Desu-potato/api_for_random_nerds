package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initial_data()
	api := gin.Default()
	api.GET("/random/mean", meanEndpoint)
	api.Run("localhost:8080")
}
