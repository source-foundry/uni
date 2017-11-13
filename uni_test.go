package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//// test version string formatting
//func TestVersionString(t *testing.T) {
//	r, _ := regexp.Compile(`\d{1,2}.\d{1,2}.\d{1,2}`)
//	if r.MatchString(version) == false {
//		t.Errorf("[FAIL] Failed to match regex pattern to version string")
//	}
//}

//// test usage string formatting
//func TestUsageString(t *testing.T) {
//	if strings.HasPrefix(usage, "Usage:") == false {
//		t.Errorf("[FAIL] Improperly formatted usage string.  Expected string to start with 'Usage:' and received %s", usage)
//	}
//}

//// test help string formatting
//func TestHelpString(t *testing.T) {
//	if strings.HasPrefix(help, "====") == false {
//		t.Errorf("[FAIL] Improperly formatted usage string. Expected to start with '===' and received %s", help)
//	}
//}

// test single argument requests to unicodeSearch function
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

// test multiple argument requests to the unicodeSearch function
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


func TestVersionString(t *testing.T) {
	r, _ := regexp.Compile(`\d{1,2}.\d{1,2}.\d{1,2}`)

	Convey("Version string formatting", t, func() {

		Convey("Given the version string constant", func() {
			v := version

			Convey("Is the version string properly formatted", func() {
				So(true, ShouldEqual, r.MatchString(v))
			})
		})
	})
}

func TestUsageString(t *testing.T) {
	Convey("Usage string formatting", t, func() {

		Convey("Given the usage string constant", func() {
			u := usage

			Convey("Does the usage string have an appropriate start substring?", func() {
				So(strings.HasPrefix(u, "Usage:"), ShouldEqual, true)
			})
		})
	})
}

func TestHelpString(t *testing.T) {
	Convey("Help string formatting", t, func() {

		Convey("Given the help string constant", func() {
			h := help

			Convey("Does the usage string have an appropriate start substring?", func() {
				So(strings.HasPrefix(h, "======="), ShouldEqual, true)
			})
		})
	})
}