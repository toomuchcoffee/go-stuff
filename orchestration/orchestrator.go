package orchestration

import (
	"example.com/gostuff/downstream"
)

func CreateResult(city string) CurrentWeatherResult {
	lonLatResponse, lonLatError := downstream.GetLonLatFromCity(city)
	weatherResponse, _ := downstream.GetWeatherForLonLat(lonLatResponse)
	interpretation := AnalyseWeather(weatherResponse.CurrentWeather.Weathercode)
	clothes := SelectClothes(weatherResponse.CurrentWeather.Temperature, interpretation.Rain)

	if lonLatError != nil {
		return CurrentWeatherResult{}
	}

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
