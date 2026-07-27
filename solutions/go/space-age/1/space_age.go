package spaceage

type Planet string

func Age(seconds float64, planet Planet) float64 {
	// Second / time * 365.25 * 24 * 60
	var time float64 = 0.0
	switch planet {
	case "Mercury":
		time = 0.2408467
	case "Venus":
		time = 0.61519726
	case "Earth":
		time = 1.0
	case "Mars":
		time = 1.8808158
	case "Jupiter":
		time = 11.862615
	case "Saturn":
		time = 29.447498
	case "Uranus":
		time = 84.016846
	case "Neptune":
		time = 164.79132
	default:
		return -1
	}
	return seconds / (time * 365.25 * 24 * 60 * 60)
}

