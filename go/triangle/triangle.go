package triangle

import "math"

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if !(a+b >= c && a+c >= b && b+c >= a) || verifySize(a) || verifySize(b) || verifySize(c) {
		k = NaT
	} else if equal(a, b) && equal(a, c) {
		k = Equ
	} else if !equal(a, b) && !equal(a, c) && !equal(b, c) {
		k = Sca
	} else {
		k = Iso
	}
	return k
}

func verifySize(size float64) bool {
	return size <= 0 || size > math.MaxFloat64
}

func equal(num1 float64, num2 float64) bool {
	return num1 == num2
}
