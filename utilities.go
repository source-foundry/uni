package main

import (
	"os"
	"strconv"
	"unicode/utf8"
)

// test os.Stdin for presence of data, if present return true, else return false
func stdinValidates(stdin *os.File) bool {
	f, err := stdin.Stat()
	if err != nil { // unable to obtain file data with Stat() method call = fail
		return false
	}

	size := f.Size()
	if size == 0 { // there does not appear to be any data in the stdin stream = fail
		return false
	}

	return true
}

func isIntInRange(value int32) bool {
	minPrintableInt, _ := strconv.ParseInt("0020", 16, 32)

	if value > utf8.MaxRune {
		// value is higher than maximum acceptable Unicode code point
		return false
	}

	if value < int32(minPrintableInt) {
		// value is lower than minimum Unicode code point for printable glyphs
		return false
	}
	// return true if integer value falls within acceptable range
	return true
}
