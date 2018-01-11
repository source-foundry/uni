package main

import (
	"strings"
	"testing"
)

// test single argument requests to unicodeSearch function for glyph --> Unicode code point search
func TestGlyphSearchCodePointsSingle(t *testing.T) {
	cases := []struct {
		unihex   string
		expected string
	}{
		{"006A", "U+006A 'j'"}, // ASCII
		{"20AC", "U+20AC '€'"}, // Currency
		{"03B2", "U+03B2 'β'"}, // Greek and Coptic
		{"0444", "U+0444 'ф'"}, // Cyrillic
		{"2580", "U+2580 '▀'"}, // Block elements
	}

	for _, c := range cases {
		response, err := glyphSearch(c.unihex)

		if err != nil {
			t.Errorf("[FAIL] Expected no error on glyph search and received %v", err)
		}

		if len(response) == 0 {
			t.Errorf("[FAIL] unicodeSearch did not return a value (len = 0)")
		} else if !strings.Contains(response, c.expected) {
			t.Errorf("[FAIL] Unicode code point '%s' yielded response %s, not expected response %s", c.unihex, response[0], c.expected)
		}
	}
}

// test single argument requests to unicodeSearch function for glyph --> Unicode code point search with appended newline
// character.  This is intended to address newline values that may come in through stdin piped text from another
// application.  Example is: `$ echo "006A" | uni -g` where newline is appended to the hexadecimal by the echo application
func TestGlyphSearchCodePointsSingleWithNewline(t *testing.T) {
	cases := []struct {
		unihex   string
		expected string
	}{
		{"006A\n", "U+006A 'j'"}, // ASCII
		{"20AC\n", "U+20AC '€'"}, // Currency
		{"03B2\n", "U+03B2 'β'"}, // Greek and Coptic
		{"0444\n", "U+0444 'ф'"}, // Cyrillic
		{"2580\n", "U+2580 '▀'"}, // Block elements
	}

	for _, c := range cases {
		response, err := glyphSearch(c.unihex)

		if err != nil {
			t.Errorf("[FAIL] Expected no error on glyph search and received %v", err)
		}

		if len(response) == 0 {
			t.Errorf("[FAIL] unicodeSearch did not return a value (len = 0)")
		} else if !strings.Contains(response, c.expected) {
			t.Errorf("[FAIL] Unicode code point '%s' yielded response %s, not expected response %s", c.unihex, response[0], c.expected)
		}
	}
}

func TestGlyphSearchInvalidHexadecimal(t *testing.T) {
	response, err := glyphSearch("zzzzz")
	if err == nil {
		t.Errorf("[FAIL] Expected error with invalid hexadecimal format and instead received response %s", response)
	}
}

func TestGlyphSearchInvalidOutOfRange(t *testing.T) {
	response, err := glyphSearch("0010")
	if err == nil {
		t.Errorf("[FAIL] Expected error with out of range hexadecimal format and instead received response %s", response)
	}
}

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
	}

	for i := range response {
		if !strings.Contains(response[i], cases[i].expected) {
			t.Errorf("[FAIL] Expected response %s but received response %s", cases[i].expected, response[i])
		}
	}
}
