package scrabble

import "strings"

func Score(word string) int {
	var score int
	for _, letter := range strings.TrimSpace(word) {
		switch strings.ToUpper(string(letter)) {
		case "Q", "Z":
			score += 10
		case "J", "X":
			score += 8
		case "K":
			score += 5
		case "F", "H", "V", "W", "Y":
			score += 4
		case "B", "C", "M", "P":
			score += 3
		case "D", "G":
			score += 2
		default:
			score += 1
		}
	}
	return score
}
