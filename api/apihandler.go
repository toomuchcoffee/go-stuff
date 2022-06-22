package api

import (
	"net/http"

	"example.com/gostuff/orchestration"
	"github.com/gin-gonic/gin"
)

func GetCurrentWeather(c *gin.Context) {
	city := c.DefaultQuery("city", orchestration.ChooseRandomCity())
	result, err := orchestration.CreateResult(city)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, orchestration.CurrentWeatherResult{Error: err.Error()})
	} else if c.Query("city") == "" {
		c.Redirect(http.StatusFound, "/weather?city="+city)
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}
