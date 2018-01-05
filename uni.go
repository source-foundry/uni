// uni is a command line executable that displays Unicode code points for glyph arguments
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	version = "0.10.0"
	usage   = "Usage: uni [glyph 1]...[glyph n]\nLine Filter Usage: [application command] | uni\n"
	help    = "=================================================\n" +
		" uni v" + version + "\n" +
		" Copyright 2018 Christopher Simpkins\n" +
		" MIT License\n\n" +
		" Source: https://github.com/source-foundry/uni\n" +
		"=================================================\n\n" +
		" Usage:\n" +
		"  - With command line arguments:\n" +
		"        $ uni [glyph 1]...[glyph n]\n" +
		"  - As line filter:\n" +
		"        $ [application command] | uni\n\n" +
		" Options:\n" +
		" -h, --help           Application help\n" +
		"     --usage          Application usage\n" +
		" -v, --version        Application version\n\n"
)

var versionShort, versionLong, helpShort, helpLong, usageLong *bool

func init() {
	// define available command line flags
	versionShort = flag.Bool("v", false, "Application version")
	versionLong = flag.Bool("version", false, "Application version")
	helpShort = flag.Bool("h", false, "Help")
	helpLong = flag.Bool("help", false, "Help")
	usageLong = flag.Bool("usage", false, "Usage")
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

	// if there are no arguments to the executable, check std input stream to see if
	// this is a line filter request that is piped to executable
	if len(os.Args) < 2 {

		if !stdinValidates(os.Stdin) {
			handleStdInErrors()
		}

		tmp := new(bytes.Buffer)
		if _, err := io.Copy(tmp, os.Stdin); err != nil {
			os.Stderr.WriteString("[Error] Failed to copy std input stream to memory. " + fmt.Sprintf("%v", err))
			os.Exit(1)
		}

		stdinList := strings.Split(tmp.String(), "") // split the stdin string by glyph to a slice
		stdOutput := unicodeSearch(stdinList)
		for _, line := range stdOutput {
			fmt.Print(line)
		}

	} else {

		// handle command line arguments to the executable
		stdOutput := unicodeSearch(os.Args[1:])
		for _, line := range stdOutput {
			fmt.Print(line)
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
	os.Stderr.WriteString("[Error] Please include at least one argument or pipe a string to the executable through the stdin stream.\n")
	os.Stderr.WriteString(usage)
	os.Exit(1)
}

// writes Unicode code point value(s) to standard output stream for glyphs entered as command line arguments
func unicodeSearch(argv []string) []string {
	var solist []string
	for i := 0; i < len(argv); i++ {
		if len(argv[i]) > 1 { // handle single argument that includes multiple glyphs
			charList := strings.Split(argv[i], "")
			for x := 0; x < len(charList); x++ {
				r, _ := utf8.DecodeRuneInString(charList[x])
				solist = append(solist, fmt.Sprintf("%#U\n", r))
			}
		} else { // handle multiple individual arguments
			r, _ := utf8.DecodeRuneInString(argv[i])
			solist = append(solist, fmt.Sprintf("%#U\n", r))
		}

	}
	return solist
}
