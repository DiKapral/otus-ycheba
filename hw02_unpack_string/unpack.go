package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(st string) (string, error) {
	strobch := ""
	flag := false
	var v rune
	for _, val := range st {
		if unicode.IsDigit(val) {

			if strobch == "" || !flag {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(val))

			if count == 0 {
				if v > 1000 {
					strobch = strobch[:len(strobch)-2]
				} else {
					strobch = strobch[:len(strobch)-1]
				}
				flag = false
				continue
			}

			count--
			str := strings.Repeat(string(v), count)

			strobch = sklei(strobch, str)
			flag = false
			continue
		} else {
			v = val
		}
		strobch = sklei(strobch, string(v))

		flag = true
	}
	return strobch, nil
}

func sklei(x, y string) string {
	var builder strings.Builder
	builder.Grow(len(x) + len(y))
	builder.WriteString(x)
	builder.WriteString(y)
	return builder.String()
}
