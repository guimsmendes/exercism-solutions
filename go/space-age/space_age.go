package space

type Planet string

const (
	SECONDS_EARTH = 31557600
	Mercury       = "Mercury"
	Venus         = "Venus"
	Earth         = "Earth"
	Mars          = "Mars"
	Jupiter       = "Jupiter"
	Saturn        = "Saturn"
	Uranus        = "Uranus"
	Neptune       = "Neptune"
)

var planets = map[Planet]float64{
	Mercury: 0.2408467,
	Venus:   0.61519726,
	Earth:   1,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132}

func Age(seconds float64, p Planet) float64 {
	return seconds / (SECONDS_EARTH * planets[p])
}
