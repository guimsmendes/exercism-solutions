package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	for _, letter := range word {

		if unicode.IsLetter(letter) && strings.Count(word, string(letter)) > 1 {
			return false
		}

	}
	return true
}
