package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func Unpack(s string) (string, error) {
	var runes []rune = []rune(s)

	var oldRune rune
	var escaped bool = false
	var result string
	var repeats int

	for _, rune := range runes { //nolint:gosimple,staticcheck
		if rune == '\\' {
			if escaped {
				escaped = false
				result += string(rune)
			} else {
				escaped = true
			}
			oldRune = rune
			continue
		}

		if unicode.IsDigit(rune) {
			if escaped {
				result += string(rune)
				oldRune = rune
				escaped = false
				continue
			}
			if oldRune == 0 {
				return "", ErrInvalidString
			}
			repeats, _ = strconv.Atoi(string(rune))
			if repeats == 0 {
				result = trimLastChar(result)
			}
			for i := 1; i < repeats; i++ {
				result += string(oldRune)
			}
			oldRune = 0
			escaped = false
			continue
		}

		if escaped {
			return "", ErrInvalidString
		}
		result += string(rune)
		oldRune = rune
		escaped = false
	}

	return result, nil
}
