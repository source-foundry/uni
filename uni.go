// uni is a command line executable that performs Unicode code point to glyph and glyph to Unicode code point searches
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	version = "1.0.0"
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
			os.Stderr.WriteString("[Error] Please include at least one argument or pipe search requests to the executable through the stdin stream.\n")
			os.Stderr.WriteString(usage)
			os.Exit(1)
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
