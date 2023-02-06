package warmup

import (
	"fmt"
	"strconv"
)

func TimeConversion(s string) string {
	// Write your code here
	hh, _ := strconv.Atoi(s[:2])
	mortem := s[len(s)-2:]
	s = s[:len(s)-2]

	var hhStr string

	if mortem == "AM" {
		if hh == 12 {
			hhStr = "00"
		} else {
			if hh < 10 {
				hhStr = fmt.Sprintf("0%d", hh)
			} else {
				hhStr = fmt.Sprintf("%d", hh)
			}
		}

	} else {
		if hh != 12 {
			hh += 12
		}
		hhStr = fmt.Sprintf("%d", hh)
	}
	s = fmt.Sprintf("%s%s", hhStr, s[2:])
	return s
}
