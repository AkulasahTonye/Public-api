package main

import (
	"github.com/Public-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/data", GetData)
	server.POST("/data", createData)

	server.Run(":3000") // Localhost
}

func GetData(context *gin.Context) {
	data := models.GetAllData()
	context.JSON(http.StatusOK, data)
}

func createData(context *gin.Context) {
	var data models.Data
	err := context.ShouldBindJSON(&data)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data."})
		return
	}

	data.Email = "akulasaht@gmail.com"
	data.GitHubUrl = "https://github.com/AkulasahTonye/API_Testing"

	data.Save()

	context.JSON(http.StatusCreated, gin.H{"Message": "Data created", "data": data})
}
