// Package weather provide weather forecasting for Goblinocus cities.
package weather

var (
// CurrentCondition stores a string and used in the weather forecasting to determine the current condition.
	CurrentCondition string
// CurrentLocation storess a string and used in the weather forecasting to determine the current location to forecast.
	CurrentLocation  string
)

// Forecast function for weather forecasting.  
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
