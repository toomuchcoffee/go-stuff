package api

import (
	"example.com/gostuff/orchestration"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentWeatherForCity(c *gin.Context) {
	city := c.Param("city")
	result, err := orchestration.CreateResult(city)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, orchestration.CurrentWeatherResult{Error: err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func GetCurrentWeather(c *gin.Context) {
	city := orchestration.ChooseRandomCity()
	c.Redirect(http.StatusFound, "/weather/"+city)
}
