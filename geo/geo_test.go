package geo

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
