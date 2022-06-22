package downstream

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CallDownstreamService[T any](url string, responseObject T) (T, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return *new(T), err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return *new(T), err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return *new(T), err
	}
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject, nil
}
