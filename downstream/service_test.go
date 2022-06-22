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

	assert.Equal(t, want, got)
}

func TestGetLonLatFromCityNotFound(t *testing.T) {
	_, gotError := GetLonLatFromCity("Foobar")
	wantError := errors.New("No LonLat found for city: Foobar")

	assert.Equal(t, wantError, gotError)
}

func TestGetLocationFromLonLat(t *testing.T) {
	lonLat := LonLat{
		Lon: "18.95586",
		Lat: "69.66558",
	}
	got := GetLocationFromLonLat(lonLat)
	want := Location{City: "TROMSØ", Country: "Norway"}

	assert.Equal(t, want, got)
}

func TestGetWeatherFromLonLat(t *testing.T) {
	lonlat := LonLat{
		Lon: "18.95586",
		Lat: "69.66558",
	}
	got, _ := GetWeatherForLonLat(lonlat)

	assert.NotNil(t, got)
}
