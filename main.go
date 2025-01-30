package main

import (
	"github.com/Public-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/", GetData)
	//server.POST("/data", createData)

	server.Run(":3000") // Localhost
}

func GetData(context *gin.Context) {
	data := models.GetAllData()
	context.JSON(http.StatusOK, data)
}
