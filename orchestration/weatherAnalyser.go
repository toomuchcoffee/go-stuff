package orchestration

func AnalyseWeather(weatherCode int) Interpretation {
	var interpretation Interpretation
	switch weatherCode {
	case 0:
		interpretation.Description = "Clear sky"
		interpretation.Rain = false
	case 1, 2, 3:
		interpretation.Description = "Mainly clear, partly cloudy, and overcast"
		interpretation.Rain = false
	case 45, 48:
		interpretation.Description = "Fog and depositing rime fog"
		interpretation.Rain = false
	case 51, 53, 55:
		interpretation.Description = "Drizzle: Light, moderate, and dense intensity"
		interpretation.Rain = true
	case 56, 57:
		interpretation.Description = "Freezing Drizzle: Light and dense intensity"
		interpretation.Rain = true
	case 61, 63, 65:
		interpretation.Description = "Rain: Slight, moderate and heavy intensity"
		interpretation.Rain = true
	case 66, 67:
		interpretation.Description = "Freezing Rain: Light and heavy intensity"
		interpretation.Rain = true
	case 71, 73, 75:
		interpretation.Description = "Snow fall: Slight, moderate, and heavy intensity"
		interpretation.Rain = true
	case 77:
		interpretation.Description = "Snow grains"
		interpretation.Rain = true
	case 80, 81, 82:
		interpretation.Description = "Rain showers: Slight, moderate, and violent"
		interpretation.Rain = true
	case 85, 86:
		interpretation.Description = "Snow showers slight and heavy"
		interpretation.Rain = true
	case 95:
		interpretation.Description = "Thunderstorm: Slight or moderate"
		interpretation.Rain = true
	case 96, 99:
		interpretation.Description = "Thunderstorm with slight and heavy hail"
		interpretation.Rain = true
	}
	return interpretation
}

type Interpretation struct {
	Description string
	Rain        bool
}
