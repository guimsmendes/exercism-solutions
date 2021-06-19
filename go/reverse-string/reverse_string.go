package reverse

func Reverse(word string) string {
	var reversed []rune
	runes := []rune(word)
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}
	return string(reversed)
}
