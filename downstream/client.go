package downstream

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallDownstreamService[T any](url string, responseObject T) T {
	client := &http.Client{}
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

	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject
}
