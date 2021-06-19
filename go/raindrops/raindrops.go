package raindrops

import "fmt"

type RainDrops struct {
	Num  int
	Name string
}

var translate = []RainDrops{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"}}

func Convert(num int) string {
	var drop string
	for _, trans := range translate {
		if num%trans.Num == 0 {
			drop += trans.Name
		}
	}
	if drop == "" {
		drop += fmt.Sprint(num)
	}

	return drop
}
