package weather

import (
	"encoding/json"
	"example.com/gostuff/geo"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWeatherForLonLat(lonlat geo.LonLat) Weather {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", lonlat.Lat, lonlat.Lon)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Weather
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject
}

type Weather struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Weathercode int     `json:"weathercode"`
	} `json:"current_weather"`
}
