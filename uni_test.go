package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
)

// Constant tests

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

func TestMainFunctionSingleGlyph(t *testing.T) {
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
	}

	if !(out == "U+006A 'j'\n") {
		t.Errorf("[FAIL] Expected execution of 'uni f' to return string U+006A 'j', but instead it returned %s", out)
	}

}

func TestMainFunctionMultiGlyph(t *testing.T) {
	os.Args = []string{"uni", "jj"}
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
	}

	if !(out == "U+006A 'j'\nU+006A 'j'\n") {
		t.Errorf("[FAIL] Expected execution of 'uni f' to return string U+006A 'j' * 2, but instead it returned %s", out)
	}

}

func TestMainFunctionSingleUnicodeCodePoint(t *testing.T) {
	os.Args = []string{"uni", "-g", "006A"}
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
	}

	if !(out == "U+006A 'j'\n") {
		t.Errorf("[FAIL] Expected execution of 'uni f' to return string U+006A 'j', but instead it returned %s", out)
	}

}

func TestMainFunctionMultiUnicodeCodePoint(t *testing.T) {
	os.Args = []string{"uni", "-g", "006A", "006A"}
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
	}

	if !(out == "U+006A 'j'\nU+006A 'j'\n") {
		t.Errorf("[FAIL] Expected execution of 'uni f' to return string U+006A 'j' * 2, but instead it returned %s", out)
	}

}
