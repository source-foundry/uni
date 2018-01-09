package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"unicode/utf8"
)

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

func TestIsIntInRangeValid(t *testing.T) {
	needle, _ := strconv.ParseInt("00ff", 16, 32)
	ok := isIntInRange(int32(needle))
	if !ok {
		t.Errorf("[FAIL] Expected hexadecimal value U+00FF to be in range and received false")
	}
}

func TestIsIntInRangeTooLarge(t *testing.T) {
	needle := utf8.MaxRune + 1
	ok := isIntInRange(needle)
	if ok {
		t.Errorf("[FAIL] Expected value `utf8.MaxRun + 1` to be out of range and received true")
	}
}

func TestIsIntInRangeTooSmall(t *testing.T) {
	needle, _ := strconv.ParseInt("0019", 16, 32)
	ok := isIntInRange(int32(needle))
	if ok {
		t.Errorf("[FAIL] Expected hexadecimal value U+0019 to be out of range and received true")
	}
}
