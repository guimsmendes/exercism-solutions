package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	var acronym string
	for _, word := range strings.Split(replace(s), " ") {
		if word != "" {
			acronym += string(word[0])
		}
	}
	return strings.ToUpper(acronym)
}

func replace(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	return strings.ReplaceAll(s, "_", "")
}
