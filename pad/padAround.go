package pad

import "strings"

func PadAround(str string, l int) string {
	x := len(str)
	if x < l {
		spaces := (l - x) / 2
		ret := strings.Repeat(" ", spaces) + str + strings.Repeat(" ", l-(spaces+x))
		return ret
	}
	return str
}
