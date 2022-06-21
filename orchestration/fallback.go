package orchestration

import (
	"math/rand"
	"time"
)

func ChooseRandomCity() string {
	rand.Seed(time.Now().Unix())
	return getCities()[rand.Intn(len(getCities()))]
}

func getCities() []string {
	cities := []string{"Berlin", "Tromso", "Reykjavik", "Longyearbyen", "Madrid", "Cape Town", "Vancouver", "Oslo", "Munich", "Rio de Janeiro", "Saskatoon"}
	return cities
}
