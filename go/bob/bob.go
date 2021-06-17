package bob

import "strings"

func Hey(remark string) string {
	switch remark = strings.TrimSpace(remark); {
	case isQuestion(remark) && isShouting(remark):
		return "Calm down, I know what I'm doing!"
	case isShouting(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	case isSilent(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func isShouting(s string) bool {
	return s == strings.ToUpper(s) && strings.ToUpper(s) != strings.ToLower(s)
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

func isSilent(s string) bool {
	return s == ""
}
