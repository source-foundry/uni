package main

import (
	"io/ioutil"
	"os"
	"testing"
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
