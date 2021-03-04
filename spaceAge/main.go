package main

import "fmt"

//Periods es una variable local tipo map
var Periods = map[string]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}
var earthOneYearSecs = 31557600.0

func main() {

	fmt.Println(age(1000000000, "Earth"))

	for planet := range Periods {
		fmt.Println(age(915170400, planet))
	}
}
func age(seconds float64, planet string) float64 {
	return seconds / (Periods[planet] * earthOneYearSecs)
}
