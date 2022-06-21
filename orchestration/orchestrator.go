package orchestration

import (
	"example.com/gostuff/downstream"
)

func CreateResult(city string) CurrentWeatherResult {
	lonLatResponse := downstream.GetLonLatFromCity(city)
	weatherResponse := downstream.GetWeatherForLonLat(lonLatResponse)
	interpretation := GetInterpretation(weatherResponse.CurrentWeather.Weathercode, weatherResponse.CurrentWeather.Temperature)

	return CurrentWeatherResult{
		City:    city,
		Lon:     lonLatResponse.Lon,
		Lat:     lonLatResponse.Lat,
		Temp:    weatherResponse.CurrentWeather.Temperature,
		Weather: interpretation.Description,
		Clothes: interpretation.Clothes,
	}
}

type CurrentWeatherResult struct {
	City    string  `json:"city"`
	Lon     string  `json:"lon"`
	Lat     string  `json:"lat"`
	Temp    float64 `json:"temp"`
	Weather string  `json:"weather"`
	Clothes string  `json:"clothes"`
}
