package main

import (
	//"fmt"
	"strings"
	"testing"
)

// tests single argument requests to unicodeSearch function
func TestUnicodeCodePointsSingle(t *testing.T) {
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

// tests multiple argument requests to the unicodeSearch function
func TestUnicodeCodePointsMultiple(t *testing.T) {
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

	var testglyphs []string

	for _, c := range cases {
		testglyphs = append(testglyphs, c.glyph)
	}

	response := unicodeSearch(testglyphs)

	if len(response) == 0 {
		t.Errorf("[FAIL] unicodeSearch did not return a value (len = 0)")
	} else if !(len(response) == 5) {
		t.Errorf("[FAIL] Expected five Unicode points in response.  Received %d", len(response))
	} else if !strings.Contains(response[0], cases[0].expected) {
		t.Errorf("[FAIL] Expected response %s for first index test but received response %s", cases[0].expected, response[0])
	}
}
