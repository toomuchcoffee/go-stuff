package weather

import (
	"example.com/gostuff/geo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLonLatFromCity(t *testing.T) {
	lonlat := geo.LonLat{
		Lon: "18.95586",
		Lat: "69.66558",
	}
	got := GetWeatherForLonLat(lonlat)

	assert.NotNil(t, got)
}
