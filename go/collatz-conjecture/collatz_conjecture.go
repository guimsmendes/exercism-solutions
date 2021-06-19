package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var count int
	if n < 1 {
		return 0, errors.New("Invalid input")
	}
	for n != 1 {
		n = calculate(n)
		count++
	}
	return count, nil
}

func calculate(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return (3 * n) + 1
	}
}
