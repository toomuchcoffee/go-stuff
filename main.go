package main

import (
	"os"

	"example.com/gostuff/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", api.GetCurrentWeather)
	router.GET("/weather", api.GetCurrentWeather)

	port := os.Getenv("PORT")
	if port == "" {
		router.Run("localhost:8080")
	} else {
		router.Run(":" + port)
	}
}
