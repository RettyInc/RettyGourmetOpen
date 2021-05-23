package business_hours

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	// 24:00未満である必要がある
	openTimeValidater = regexp.MustCompile(`^([0-1]?\d|2[0-3]):[0-5]\d$`)
	// 終了時間は24:00を超える表記を許容する
	closeTimeValidater = regexp.MustCompile(`^[0-2]?\d:[0-5]\d$`)
)

func IsValidBusinessHours(hours []string) bool {
	var lastClose string

	for i := 0; i < len(hours); i = i + 2 {
		open, close := hours[i], hours[i+1]

		if open == "" && close == "" {
			break
		}

		ret := isValidBusinessHourPair(open, close)
		if ret == false {
			return false
		}

		if lastClose != "" {
			ret := isValidBusinessHourOrder(lastClose, open)
			if ret == false {
				return false
			}
		}

		lastClose = close
	}

	return true
}

func isValidBusinessHourPair(open, close string) bool {
	ret := true

	if open == "" {
		ret = closeTimeValidater.MatchString(close)
	} else if close == "" {
		ret = openTimeValidater.MatchString(open)
	} else if open != "" && close != "" {
		ret = openTimeValidater.MatchString(open) && closeTimeValidater.MatchString(close) && isValidBusinessHourOrder(open, close)
	}

	return ret
}

func isValidBusinessHourOrder(first, second string) bool {
	firstVal, _ := strconv.Atoi(strings.Replace(first, ":", "", -1))
	secondVal, _ := strconv.Atoi(strings.Replace(second, ":", "", -1))

	return firstVal < secondVal
}
