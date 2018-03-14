package space

// Planet represent the names of planets in the solar system.
type Planet string

const secondsInEarthYear = 31557600

var orbitalPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the relative age of a person on a given planet.
func Age(age float64, p Planet) float64 {
	return age / (secondsInEarthYear * orbitalPeriods[p])
}
