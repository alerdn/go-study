package util

import "strings"

func BetweenPad(strLeft string, strRight string, totalLength int, pad string) string {
	strLeftLen := len(strLeft)
	strRightLen := len(strRight)

	lenCombined := strLeftLen + strRightLen
	if lenCombined >= totalLength {
		return strLeft + " " + strRight
	}

	spaceBetween := totalLength - lenCombined
	return strLeft + strings.Repeat(pad, spaceBetween) + strRight
}
