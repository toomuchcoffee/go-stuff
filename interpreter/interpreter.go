package interpreter

func AnalyseWeather(weatherCode int, temperature float64) Interpretation {
	interpretation := analyseWeather(weatherCode, temperature)

	interpretation.Clothes = selectClothes(interpretation.Temp, interpretation.Rain)

	return interpretation
}

func selectClothes(temperature float64, rain bool) string {
	var clothes string

	if temperature > 20 {
		clothes += "light "
	} else if temperature < 0 {
		clothes += "heavy "
	} else {
		clothes += "medium "
	}

	if rain {
		clothes += "waterproof "
	}
	clothes += "wear"
	return clothes
}

func analyseWeather(weatherCode int, temperature float64) Interpretation {
	var mappedWeather = Interpretation{Temp: temperature}
	switch weatherCode {
	case 0:
		mappedWeather.Description = "Clear sky"
		mappedWeather.Rain = false
	case 1, 2, 3:
		mappedWeather.Description = "Mainly clear, partly cloudy, and overcast"
		mappedWeather.Rain = false
	case 45, 48:
		mappedWeather.Description = "Fog and depositing rime fog"
		mappedWeather.Rain = false
	case 51, 53, 55:
		mappedWeather.Description = "Drizzle: Light, moderate, and dense intensity"
		mappedWeather.Rain = true
	case 56, 57:
		mappedWeather.Description = "Freezing Drizzle: Light and dense intensity"
		mappedWeather.Rain = true
	case 61, 63, 65:
		mappedWeather.Description = "Rain: Slight, moderate and heavy intensity"
		mappedWeather.Rain = true
	case 66, 67:
		mappedWeather.Description = "Freezing Rain: Light and heavy intensity"
		mappedWeather.Rain = true
	case 71, 73, 75:
		mappedWeather.Description = "Snow fall: Slight, moderate, and heavy intensity"
		mappedWeather.Rain = true
	case 77:
		mappedWeather.Description = "Snow grains"
		mappedWeather.Rain = true
	case 80, 81, 82:
		mappedWeather.Description = "Rain showers: Slight, moderate, and violent"
		mappedWeather.Rain = true
	case 85, 86:
		mappedWeather.Description = "Snow showers slight and heavy"
		mappedWeather.Rain = true
	case 95:
		mappedWeather.Description = "Thunderstorm: Slight or moderate"
		mappedWeather.Rain = true
	case 96, 99:
		mappedWeather.Description = "Thunderstorm with slight and heavy hail"
		mappedWeather.Rain = true
	}
	return mappedWeather
}

type Interpretation struct {
	Description string
	Rain        bool
	Temp        float64
	Clothes     string
}
