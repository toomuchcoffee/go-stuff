package orchestration

import (
	"testing"
)

func TestSelectClothes(t *testing.T) {
	type args struct {
		temperature float64
		rain        bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "light wear", args: args{temperature: 21, rain: false}, want: "light wear"},
		{name: "medium wear", args: args{temperature: 11, rain: false}, want: "medium wear"},
		{name: "heavy wear", args: args{temperature: -1, rain: false}, want: "heavy wear"},
		{name: "light waterproof wear", args: args{temperature: 21, rain: true}, want: "light waterproof wear"},
		{name: "medium waterproof wear", args: args{temperature: 11, rain: true}, want: "medium waterproof wear"},
		{name: "heavy waterproof wear", args: args{temperature: -1, rain: true}, want: "heavy waterproof wear"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectClothes(tt.args.temperature, tt.args.rain); got != tt.want {
				t.Errorf("SelectClothes() = %v, want %v", got, tt.want)
			}
		})
	}
}
