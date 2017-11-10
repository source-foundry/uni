package main

import (
	"strings"
	"testing"
)

func TestUnicodeCodePoints(t *testing.T) {
	cases := []struct {
		glyph    string
		expected string
	}{
		{"j", "U+006A 'j'\n"}, // ASCII
		{"€", "U+20AC '€'\n"}, // Currency
		{"β", "U+03B2 'β'\n"}, // Greek and Coptic
		{"ф", "U+0444 'ф'\n"}, // Cyrillic
		{"▀", "U+2580 '▀'\n"}, // Block elements
	}

	for _, c := range cases {
		testglyph := []string{c.glyph}
		response := unicodeSearch(testglyph)

		if len(response) == 0 {
			t.Errorf("[FAIL] unicodeSearch did not return a value (len = 0)")
		} else if !strings.Contains(response[0], c.expected) {
			t.Errorf("[FAIL] Glyph '%s' yielded response %s, not expected response %s", c.glyph, response[0], c.expected)
		}
	}
}
