package downstream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLonLatFromCity(t *testing.T) {
	got := GetLonLatFromCity("Tromso")
	want := LonLat{
		Lon: "18.95586",
		Lat: "69.66558",
	}

	assert.Equal(t, got, want)
}

func TestGetWeatherFromLonLat(t *testing.T) {
	lonlat := LonLat{
		Lon: "18.95586",
		Lat: "69.66558",
	}
	got := GetWeatherForLonLat(lonlat)

	assert.NotNil(t, got)
}
