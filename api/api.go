package api

import (
	"example.com/gostuff/geo"
	"example.com/gostuff/interpreter"
	"example.com/gostuff/weather"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func GetCurrentWeatherForCity(c *gin.Context) {
	city := c.Param("city")
	result := createResult(city)
	c.IndentedJSON(http.StatusOK, result)
}

func GetCurrentWeather(c *gin.Context) {
	city := chooseRandomCity()
	c.Redirect(http.StatusFound, "/weather/"+city)
}

type CurrentWeatherResult struct {
	City    string  `json:"city"`
	Lon     string  `json:"lon"`
	Lat     string  `json:"lat"`
	Temp    float64 `json:"temp"`
	Weather string  `json:"weather"`
	Clothes string  `json:"clothes"`
}

func createResult(city string) CurrentWeatherResult {
	lonLatResponse := geo.GetLonLatFromCity(city)
	weatherResponse := weather.GetWeatherForLonLat(lonLatResponse)
	interpretation := interpreter.AnalyseWeather(weatherResponse.CurrentWeather.Weathercode, weatherResponse.CurrentWeather.Temperature)

	return CurrentWeatherResult{
		City:    city,
		Lon:     lonLatResponse.Lon,
		Lat:     lonLatResponse.Lat,
		Temp:    interpretation.Temp,
		Weather: interpretation.Description,
		Clothes: interpretation.Clothes,
	}
}

func chooseRandomCity() string {
	cities := []string{"Berlin", "Tromso", "Reykjavik", "Longyearbyen", "Madrid", "Cape Town", "Vancouver", "Oslo", "Munich", "Rio de Janeiro", "Saskatoon"}
	rand.Seed(time.Now().Unix())
	city := cities[rand.Intn(len(cities))]
	return city
}
