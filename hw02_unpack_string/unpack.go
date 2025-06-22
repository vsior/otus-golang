package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const (
	ignoreR = rune(-1)
	escapeR = rune('\\')
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	runes := []rune(s)
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	prevR := ignoreR
	var escaped bool
	var err error
	b := strings.Builder{}

	for _, r := range runes {
		if prevR, escaped, err = processRune(r, prevR, escaped, &b); err != nil {
			return "", err
		}
	}

	if prevR == escapeR && !escaped {
		return "", ErrInvalidString
	}

	if !unicode.IsDigit(prevR) || escaped {
		b.WriteRune(prevR)
	}

	return b.String(), nil
}

func processRune(cur rune, prev rune, escaped bool, builder *strings.Builder) (rune, bool, error) {
	if unicode.IsDigit(prev) && !escaped && unicode.IsDigit(cur) {
		return cur, false, ErrInvalidString
	}
	// FIX
	if prev == escapeR && !escaped && (cur != escapeR && !unicode.IsDigit(cur)) {
		return cur, false, ErrInvalidString
	}

	if !escaped && unicode.IsDigit(prev) {
		prev = ignoreR
	}

	escaped = !escaped && prev == escapeR
	if escaped && cur == escapeR {
		prev = ignoreR
	}

	if unicode.IsDigit(cur) {
		if escaped {
			return cur, escaped, nil
		}
		count, err := strconv.Atoi(string(cur))
		if err != nil {
			return rune(0), false, err
		}
		builder.WriteString(strings.Repeat(string(prev), count))
		return cur, escaped, nil
	}

	if prev != ignoreR {
		builder.WriteRune(prev)
	}

	return cur, escaped, nil
}
