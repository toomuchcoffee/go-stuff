package main

import (
	"example.com/gostuff/apiservice"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/weather/:city", getCurrentWeather)

	port := os.Getenv("PORT")
	if port == "" {
		router.Run("localhost:8080")
	} else {
		router.Run(":" + port)
	}
}

func getCurrentWeather(c *gin.Context) {
	city := c.Param("city")
	result := apiservice.SuggestClothes(city)
	c.IndentedJSON(http.StatusOK, result)
}
