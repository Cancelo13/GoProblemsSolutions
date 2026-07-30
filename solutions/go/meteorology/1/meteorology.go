package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (te TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[te]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (te Temperature) String() string {
	return fmt.Sprintf("%v %s", te.degree, te.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (s SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[s]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%v %s", s.magnitude, s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (data MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %v%s Humidity", data.location, data.temperature.String(), data.windDirection, data.windSpeed.String(), data.humidity, "%")
}