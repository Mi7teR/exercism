// Package weather contains Forecast method that returns string representing weather condition for given location.
package weather

// CurrentCondition - current weather condition.
var CurrentCondition string

// CurrentLocation - current location.
var CurrentLocation string

// Forecast - return string that representing given current weather condition with given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
