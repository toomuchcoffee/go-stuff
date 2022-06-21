package apiservice

import (
	"example.com/gostuff/geo"
	"example.com/gostuff/interpreter"
	"example.com/gostuff/weather"
)

func SuggestClothes(city string) Suggestion {

	lonLatResponse := geo.GetLonLatFromCity(city)
	weatherResponse := weather.GetWeatherForLonLat(lonLatResponse)
	interpretation := interpreter.AnalyseWeather(weatherResponse.CurrentWeather.Weathercode, weatherResponse.CurrentWeather.Temperature)

	return Suggestion{
		City:    city,
		Lon:     lonLatResponse.Lon,
		Lat:     lonLatResponse.Lat,
		Temp:    interpretation.Temp,
		Weather: interpretation.Description,
		Clothes: interpretation.Clothes,
	}
}

type Suggestion struct {
	City    string  `json:"city"`
	Lon     string  `json:"lon"`
	Lat     string  `json:"lat"`
	Temp    float64 `json:"temp"`
	Weather string  `json:"weather"`
	Clothes string  `json:"clothes"`
}
