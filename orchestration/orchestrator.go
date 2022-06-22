package orchestration

import (
	"example.com/gostuff/downstream"
)

func CreateResult(city string) (CurrentWeatherResult, error) {
	lonLatResponse, lonLatError := downstream.GetLonLatFromCity(city)
	if lonLatError != nil {
		return CurrentWeatherResult{}, lonLatError
	}
	weatherResponse, _ := downstream.GetWeatherForLonLat(lonLatResponse)
	interpretation := AnalyseWeather(weatherResponse.CurrentWeather.Weathercode)
	clothes := SelectClothes(weatherResponse.CurrentWeather.Temperature, interpretation.Rain)

	return CurrentWeatherResult{
		City:    city,
		Lon:     lonLatResponse.Lon,
		Lat:     lonLatResponse.Lat,
		Temp:    weatherResponse.CurrentWeather.Temperature,
		Weather: interpretation.Description,
		Clothes: clothes,
	}, nil
}

type CurrentWeatherResult struct {
	City    string  `json:"city,omitempty"`
	Lon     string  `json:"lon,omitempty"`
	Lat     string  `json:"lat,omitempty"`
	Temp    float64 `json:"temp,omitempty"`
	Weather string  `json:"weather,omitempty"`
	Clothes string  `json:"clothes,omitempty"`
	Error   string  `json:"error,omitempty"`
}
