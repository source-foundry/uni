package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// glyphSearch identifies glyphs for Unicode code point (hexadecimal string) search requests
func glyphSearch(arg string) (string, error) {
	i, err := strconv.ParseInt(arg, 16, 32)
	if err != nil {
		return "", err
	}
	// test to confirm that Unicode code point falls in range U+0020 = min printable glyph to utf8.MaxRune value
	ok := isIntInRange(int32(i))
	if !ok {
		errmsg := fmt.Errorf("hexadecimal value '%s' was out of range", arg)
		return "", errmsg
	}
	r := rune(i)
	stdoutString := "U+" + arg + " '" + string(r) + "'"
	return stdoutString, nil
}

// unicodeSearch identifies the Unicode code point for glyph search requests
func unicodeSearch(argv []string) []string {
	var solist []string
	for i := 0; i < len(argv); i++ {
		if len(argv[i]) > 1 { // handle single argument that includes multiple glyphs
			charList := strings.Split(argv[i], "")
			for x := 0; x < len(charList); x++ {
				r, _ := utf8.DecodeRuneInString(charList[x])
				solist = append(solist, fmt.Sprintf("%#U", r))
			}
		} else { // handle multiple individual arguments
			r, _ := utf8.DecodeRuneInString(argv[i])
			solist = append(solist, fmt.Sprintf("%#U", r))
		}

	}
	return solist
}
