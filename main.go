package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type data struct {
	// struct tag
	Email           string    `json:"email"`
	CurrentDatetime time.Time `json:"current_datetime"`
	GithubUrl       string    `json:"github_url"`
}

func main() {
	server := gin.Default()
	getAllData()

	server.GET("data", GetData)

	server.Run("localhost:8080") // Localhost
}

func GetData(context *gin.Context) {
	data := getAllData()
	context.JSON(http.StatusOK, data)
}

func getAllData() data {

	currentTime := time.Now().UTC()

	data := data{
		Email:           "akulasaht@gmail.com",
		CurrentDatetime: currentTime,
		GithubUrl:       "https://github.com/AkulasahTonye/API_Testing",
	}
	return data
}
