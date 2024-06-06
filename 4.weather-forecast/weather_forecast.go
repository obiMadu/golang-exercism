// Package weather forecasts weather conditions for the country of Goblinocus
package weather
// CurrentCondition variable stores the current weather condition

// CurrentLocation variable stores the current location

var CurrentCondition string
var CurrentLocation string

// Forecast takes in a city and condition parameter and returns a forecast
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
