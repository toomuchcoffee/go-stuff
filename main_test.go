package main

import (
	"encoding/json"
	"example.com/gostuff/apiservice"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetCurrentWeather(t *testing.T) {
	r := SetUpRouter()
	r.GET("/weather/:city", getCurrentWeather)
	req, _ := http.NewRequest("GET", "/weather/Tromso", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var got apiservice.Suggestion
	json.Unmarshal(w.Body.Bytes(), &got)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, got.Weather)
	assert.NotNil(t, got.Temp)
	assert.NotNil(t, got.Lat)
	assert.NotNil(t, got.Lon)
	assert.Equal(t, "18.95586", got.Lon)
	assert.Equal(t, "69.66558", got.Lat)
	assert.Equal(t, "Tromso", got.City)
	assert.NotNil(t, got.Clothes)
}
