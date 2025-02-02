package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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

	server.GET("/", GetData)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}

	// Bind to 0.0.0.0 to allow external access
	server.Run("0.0.0.0:" + port)
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
