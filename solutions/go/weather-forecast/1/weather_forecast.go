// Package weather provides weather information for a location,
// including the current weather condition and a formatted forecast.
package weather

var (
	// CurrentCondition stores the current weather condition for the
	// most recently requested location.
	CurrentCondition string

	// CurrentLocation stores the name of the location for the
	// most recently requested forecast.
	CurrentLocation string
)

// Forecast returns a formatted forecast string for the specified city
// and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}