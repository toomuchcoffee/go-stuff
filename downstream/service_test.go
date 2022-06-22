package downstream

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLonLatFromCity(t *testing.T) {
	got, _ := GetLonLatFromCity("Tromso")
	want := LonLat{
		Lon: "18.95586",
		Lat: "69.66558",
	}

	assert.Equal(t, got, want)
}

func TestGetLonLatFromCityNotFound(t *testing.T) {
	_, gotError := GetLonLatFromCity("Foobar")
	wantError := errors.New("No LonLat found for city: Foobar")

	assert.Equal(t, wantError, gotError)
}

func TestGetWeatherFromLonLat(t *testing.T) {
	lonlat := LonLat{
		Lon: "18.95586",
		Lat: "69.66558",
	}
	got, _ := GetWeatherForLonLat(lonlat)

	assert.NotNil(t, got)
}
