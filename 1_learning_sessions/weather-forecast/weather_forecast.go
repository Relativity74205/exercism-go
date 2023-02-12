// Package weather provides the forecast for a city.
package weather

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation describes the current weather location.
var CurrentLocation string

// Forecast returns a string representation the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
