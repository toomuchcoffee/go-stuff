package orchestration

import "strings"

func SelectClothes(temperature float64, rain bool) string {
	var clothes []string
	if temperature >= 20 {
		clothes = append(clothes, "light")
	} else if temperature <= 0 {
		clothes = append(clothes, "heavy")
	} else {
		clothes = append(clothes, "medium")
	}
	if rain {
		clothes = append(clothes, "waterproof")
	}
	clothes = append(clothes, "wear")
	return strings.Join(clothes, " ")
}
