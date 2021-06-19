package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(input string) bool {
	sum := 0
	inputArr := strings.Split(strings.ReplaceAll(input, " ", ""), "")

	if len(inputArr) <= 1 {
		return false
	}

	for i, number := range inputArr {
		if !unicode.IsDigit(rune(number[0])) {
			return false
		}
		v, _ := strconv.Atoi(number)
		if len(inputArr)%2 == i%2 {
			v = luhnAlgorithm(v)
		}
		sum += v
	}
	return sum%10 == 0
}

func luhnAlgorithm(v int) int {
	v *= 2
	if v > 9 {
		v -= 9
	}
	return v
}
