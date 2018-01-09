package main

import (
	"strings"
	"testing"
)

// test single argument requests to unicodeSearch function for glyph --> Unicode code point search
func TestUnicodeSearchCodePointsSingle(t *testing.T) {
	cases := []struct {
		glyph    string
		expected string
	}{
		{"j", "U+006A 'j'"}, // ASCII
		{"€", "U+20AC '€'"}, // Currency
		{"β", "U+03B2 'β'"}, // Greek and Coptic
		{"ф", "U+0444 'ф'"}, // Cyrillic
		{"▀", "U+2580 '▀'"}, // Block elements
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

// test multiple argument requests to the unicodeSearch function for glyph --> Unicode code point search
func TestUnicodeSearchCodePointsMultiple(t *testing.T) {
	cases := []struct {
		glyph    string
		expected string
	}{
		{"j", "U+006A 'j'"}, // ASCII
		{"€", "U+20AC '€'"}, // Currency
		{"β", "U+03B2 'β'"}, // Greek and Coptic
		{"ф", "U+0444 'ф'"}, // Cyrillic
		{"▀", "U+2580 '▀'"}, // Block elements
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
