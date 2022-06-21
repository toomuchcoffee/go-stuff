package interpreter

import (
	"testing"
)

func TestAnalyseWeather(t *testing.T) {
	type args struct {
		weatherCode int
		temperature float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "light wear", args: args{weatherCode: 3, temperature: 21}, want: "light wear"},
		{name: "medium wear", args: args{weatherCode: 3, temperature: 11}, want: "medium wear"},
		{name: "heavy wear", args: args{weatherCode: 3, temperature: -1}, want: "heavy wear"},
		{name: "light waterproof wear", args: args{weatherCode: 55, temperature: 21}, want: "light waterproof wear"},
		{name: "medium waterproof wear", args: args{weatherCode: 55, temperature: 11}, want: "medium waterproof wear"},
		{name: "heavy waterproof wear", args: args{weatherCode: 55, temperature: -1}, want: "heavy waterproof wear"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnalyseWeather(tt.args.weatherCode, tt.args.temperature); got.Clothes != tt.want {
				t.Errorf("AnalyseWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}
