package api

import (
	"example.com/gostuff/orchestration"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentWeatherForCity(c *gin.Context) {
	city := c.Param("city")
	result := orchestration.CreateResult(city)
	c.IndentedJSON(http.StatusOK, result)
}

func GetCurrentWeather(c *gin.Context) {
	city := orchestration.ChooseRandomCity()
	c.Redirect(http.StatusFound, "/weather/"+city)
}
