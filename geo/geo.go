package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetLonLatFromCity(city string) LonLat {
	client := &http.Client{}
	url := fmt.Sprintf("https://geocode.xyz/%s?json=1", city)
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
	var responseObject LonLat
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject
}

type LonLat struct {
	Lon string `json:"longt"`
	Lat string `json:"latt"`
}
