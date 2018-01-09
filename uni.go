// uni is a command line executable that performs Unicode code point to glyph and glyph to Unicode code point searches
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	version = "0.11.0"
	usage   = "Usage: uni (options) [arg 1]...[arg n]\nLine Filter Usage: [application command] | uni (options)\n"
	help    = "=================================================\n" +
		" uni v" + version + "\n" +
		" Copyright 2018 Christopher Simpkins\n" +
		" MIT License\n\n" +
		" Source: https://github.com/source-foundry/uni\n" +
		"=================================================\n\n" +
		" Usage:\n" +
		"  - With command line arguments:\n" +
		"        $ uni (options) [arg 1]...[arg n]\n" +
		"  - As line filter:\n" +
		"        $ [application command] | uni (options)\n\n" +
		" Options:\n" +
		" -g, --glyph          Search for glyph with Unicode code point\n" +
		" -h, --help           Application help\n" +
		"     --usage          Application usage\n" +
		" -v, --version        Application version\n\n"
)

var versionShort, versionLong, helpShort, helpLong, usageLong, glyphShort, glyphLong *bool
var isGlyphSearch = false

func init() {
	// define available command line flags
	versionShort = flag.Bool("v", false, "Application version")
	versionLong = flag.Bool("version", false, "Application version")
	helpShort = flag.Bool("h", false, "Help")
	helpLong = flag.Bool("help", false, "Help")
	usageLong = flag.Bool("usage", false, "Usage")
	glyphShort = flag.Bool("g", false, "Glyph")
	glyphLong = flag.Bool("glyph", false, "Glyph")
}

func main() {

	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		os.Stdout.WriteString("uni v" + version + "\n")
		os.Exit(0)
	case *helpShort, *helpLong:
		os.Stdout.WriteString(help)
		os.Exit(0)
	case *usageLong:
		os.Stdout.WriteString(usage)
		os.Exit(0)
	}

	// create a flag for Unicode code point to glyph search command line requests
	if *glyphShort || *glyphLong {
		isGlyphSearch = true
	}

	// create a flag for stdin stream search based upon number of arguments and use of glyph search flag
	isStdinSearch := (!isGlyphSearch && len(os.Args) < 2) || (isGlyphSearch && len(os.Args) < 3)

	if isStdinSearch {

		if !stdinValidates(os.Stdin) {
			handleStdInErrors()
		}

		tmp := new(bytes.Buffer)
		if _, err := io.Copy(tmp, os.Stdin); err != nil {
			os.Stderr.WriteString("[Error] Failed to copy std input stream to memory. " + fmt.Sprintf("%v", err))
			os.Exit(1)
		}

		// stdin stream search for glyph from Unicode code point search request
		if isGlyphSearch {
			stdinList := strings.Split(tmp.String(), ` `)
			for _, arg := range stdinList {
				stdoutString, err := glyphSearch(arg)
				if err != nil {
					errmsg := fmt.Sprintf("%v", err)
					os.Stderr.WriteString("[Error] Unable to parse the Unicode code point request '" + arg + "' to a glyph. " + errmsg + "\n")
					os.Exit(1)
				}
				fmt.Println(stdoutString)
			}

		} else { // stdin stream search for Unicode code point from glyph search request
			stdinList := strings.Split(tmp.String(), "") // split the stdin string by glyph to a slice
			stdOutputList := unicodeSearch(stdinList)
			for _, stdoutString := range stdOutputList {
				fmt.Println(stdoutString)
			}
		}

	} else {
		// argument search for glyph from Unicode code point search request
		if isGlyphSearch {
			for _, arg := range os.Args[2:] {
				stdoutString, err := glyphSearch(arg)
				if err != nil {
					errmsg := fmt.Sprintf("%v", err)
					os.Stderr.WriteString("[Error] Unable to parse the Unicode code point request '" + arg + "' to a glyph. " + errmsg + "\n")
					os.Exit(1)
				}
				fmt.Println(stdoutString)
			}
		} else { // argument search for Unicode code point from glyph search request
			stdOutputList := unicodeSearch(os.Args[1:])
			for _, stdoutString := range stdOutputList {
				fmt.Println(stdoutString)
			}
		}

	}

}

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

func handleStdInErrors() {
	os.Stderr.WriteString("[Error] Please include at least one argument or pipe search requests to the executable through the stdin stream.\n")
	os.Stderr.WriteString(usage)
	os.Exit(1)
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

// glyphSearch identifies glyphs for Unicode code point (hexadecimal string) search requests
func glyphSearch(arg string) (string, error) {
	i, err := strconv.ParseInt(arg, 16, 32)
	if err != nil {
		return "", err
	}
	// test to confirm that Unicode code point falls in range U+0020 = min printable glyph to utf8.MaxRune value
	ok := isIntInRange(int32(i))
	if !ok {
		errmsg := fmt.Errorf("Hexadecimal value '" + arg + "' was out of range.")
		return "", errmsg
	}
	r := rune(i)
	stdoutString := "U+" + arg + " '" + string(r) + "'"
	return stdoutString, nil
}

// unicodeSearch identifies the Unicode code point for glyph search requests
func unicodeSearch(argv []string) []string {
	var solist []string
	for i := 0; i < len(argv); i++ {
		if len(argv[i]) > 1 { // handle single argument that includes multiple glyphs
			charList := strings.Split(argv[i], "")
			for x := 0; x < len(charList); x++ {
				r, _ := utf8.DecodeRuneInString(charList[x])
				solist = append(solist, fmt.Sprintf("%#U", r))
			}
		} else { // handle multiple individual arguments
			r, _ := utf8.DecodeRuneInString(argv[i])
			solist = append(solist, fmt.Sprintf("%#U", r))
		}

	}
	return solist
}
