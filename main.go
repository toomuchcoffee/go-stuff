package main

import (
	"example.com/gostuff/api"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/", api.GetCurrentWeather)
	router.GET("/weather", api.GetCurrentWeather)
	router.GET("/weather/:city", api.GetCurrentWeatherForCity)

	port := os.Getenv("PORT")
	if port == "" {
		router.Run("localhost:8080")
	} else {
		router.Run(":" + port)
	}
}
