package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
)

// test single argument requests to unicodeSearch function
func TestUnicodeCodePointsSingle(t *testing.T) {
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

// test multiple argument requests to the unicodeSearch function
func TestUnicodeCodePointsMultiple(t *testing.T) {
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

func TestMainFunction(t *testing.T) {
	os.Args = []string{"uni", "j"}
	old := os.Stdout // keep backup of stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w

	outC := make(chan string)

	// copy the output in a separate goroutine
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	main() // call main function with mock os.Args defined above

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if len(out) == 0 {
		t.Errorf("[FAIL] Test of main() function did not return standard output response")
	} else if !strings.HasPrefix(out, "U+006A") {
		t.Errorf("[FAIL] Expected execution of 'uni f' to return string that begins with 'U+006A', but instead it returned %s", out)
	}

}

func TestStdinValidatesTrueFunction(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "stdin")
	defer os.Remove(file.Name())

	file.WriteString("stdin test")

	result := stdinValidates(file)
	if result != true {
		t.Errorf("[FAIL] Attempt to validate mocked stdin failed.")
	}

}

func TestStdinValidatesFalseFunction(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "stdin")
	defer os.Remove(file.Name())

	file.WriteString("")

	result := stdinValidates(file)
	if result != false {
		t.Errorf("[FAIL] Attempt to validate empty mocked stdin failed.")
	}

}

func TestVersionString(t *testing.T) {
	r, _ := regexp.Compile(`\d{1,2}.\d{1,2}.\d{1,2}`)
	// match expected format of version string
	if r.MatchString(version) != true {
		t.Errorf("[FAIL] Unexpected version string format identified.")
	}
}

func TestUsageString(t *testing.T) {
	if strings.HasPrefix(usage, "Usage:") == false {
		t.Errorf("[FAIL] Unexpected usage string format.")
	}
}

func TestHelpString(t *testing.T) {
	if strings.HasPrefix(help, "=====") == false {
		t.Errorf("[FAIL] Unexpected help string format.")
	}
}
