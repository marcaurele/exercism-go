// Package weather provides weather tools.
package weather

// CurrentCondition represents the current condition.
var CurrentCondition string

// CurrentLocation represents a location.
var CurrentLocation string

// Forecast returns a string for the weather forecast in city for the condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
