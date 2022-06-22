package orchestration

import (
	"example.com/gostuff/downstream"
)

func CreateResult(city string) CurrentWeatherResult {
	lonLatResponse := downstream.GetLonLatFromCity(city)
	weatherResponse := downstream.GetWeatherForLonLat(lonLatResponse)
	interpretation := AnalyseWeather(weatherResponse.CurrentWeather.Weathercode)
	clothes := SelectClothes(weatherResponse.CurrentWeather.Temperature, interpretation.Rain)

	return CurrentWeatherResult{
		City:    city,
		Lon:     lonLatResponse.Lon,
		Lat:     lonLatResponse.Lat,
		Temp:    weatherResponse.CurrentWeather.Temperature,
		Weather: interpretation.Description,
		Clothes: clothes,
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
