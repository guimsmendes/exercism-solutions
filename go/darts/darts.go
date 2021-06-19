package darts

import "math"

const (
	inner = 1
	mid   = 5
	outer = 10
	low   = 1
	high  = 10
)

func Score(x float64, y float64) int {
	rad := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	switch {
	case rad <= inner:
		return high
	case rad <= mid:
		return mid
	case rad <= outer:
		return low
	default:
		return 0
	}
}
