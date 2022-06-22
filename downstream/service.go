package downstream

import (
	"errors"
	"fmt"
)

func GetLonLatFromCity(city string) (LonLat, error) {
	url := fmt.Sprintf("https://geocode.xyz/%s?json=1", city)
	var responseObject LonLat
	lonLat, _ := CallDownstreamService(url, responseObject)
	if lonLat.Lat == "" || lonLat.Lon == "" {
		return lonLat, errors.New("No LonLat found for city: " + city)
	}
	return lonLat, nil
}

type LonLat struct {
	Lon string `json:"longt"`
	Lat string `json:"latt"`
}

func GetWeatherForLonLat(lonlat LonLat) (Weather, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", lonlat.Lat, lonlat.Lon)
	var responseObject Weather
	return CallDownstreamService(url, responseObject)
}

type Weather struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Weathercode int     `json:"weathercode"`
	} `json:"current_weather"`
}
