package util

import "strings"

func CenterPad(str string, totalLength int, pad string) string {
	strLen := len(str)
	if strLen >= totalLength {
		return str
	}

	padding := totalLength - strLen
	left := padding / 2
	right := padding - left

	return strings.Repeat(pad, left) + str + strings.Repeat(pad, right)
}
