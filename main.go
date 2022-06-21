package main

import (
	"example.com/gostuff/apiservice"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/weather/:city", getCurrentWeather)

	router.Run("localhost:8080")
}

func getCurrentWeather(c *gin.Context) {
	city := c.Param("city")
	result := apiservice.SuggestClothes(city)
	c.IndentedJSON(http.StatusOK, result)
}
