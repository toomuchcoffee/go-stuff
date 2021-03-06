package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/gostuff/api"
	"example.com/gostuff/orchestration"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetCurrentWeatherForCity(t *testing.T) {
	r := SetUpRouter()
	r.GET("/weather", api.GetCurrentWeather)
	req, _ := http.NewRequest("GET", "/weather?city=Tromso", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var got orchestration.CurrentWeatherResult
	json.Unmarshal(w.Body.Bytes(), &got)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, got.Weather)
	assert.NotNil(t, got.Temp)
	assert.NotNil(t, got.Lat)
	assert.NotNil(t, got.Lon)
	assert.Equal(t, "18.95586", got.Lon)
	assert.Equal(t, "69.66558", got.Lat)
	assert.Equal(t, "Norway", got.Country)
	assert.Equal(t, "TROMSØ", got.City)
	assert.NotNil(t, got.Clothes)
}
