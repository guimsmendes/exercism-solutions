package rotationalcipher

import "unicode"

var alpha string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RotationalCipher(plain string, shiftKey int) string {
	a := []rune(alpha)
	var ret []rune
	for _, letter := range plain {
		p, b := contains(letter)
		if b {
			if letter == a[p] {
				ret = append(ret, a[getPos(p, shiftKey)])
			} else {
				ret = append(ret, unicode.ToLower(a[getPos(p, shiftKey)]))
			}
		} else {
			ret = append(ret, letter)
		}
	}
	return string(ret)
}
func contains(letter rune) (int, bool) {
	for i, l := range alpha {
		if l == unicode.ToUpper(letter) {
			return i, true
		}
	}
	return 0, false
}

func getPos(p int, shiftKey int) int {
	if p+shiftKey >= len(alpha) {
		return p + shiftKey - len(alpha)
	}
	return p + shiftKey
}
