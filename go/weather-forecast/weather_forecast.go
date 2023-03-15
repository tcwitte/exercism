// Package weather provides functionality for forecasting the weather.
package weather

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation represents the location for which a weather forecast is done.
var CurrentLocation string

// Forecast constructs a weather forecast from the given city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
