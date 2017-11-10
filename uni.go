// uni is a command line executable that displays Unicode code points for glyph arguments
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	VERSION = "0.9.0"
	USAGE   = "Usage: uni [glyph 1]...[glyph n]\n"
	HELP    = "=================================================\n" +
		" uni v" + VERSION + "\n" +
		" Copyright 2017 Christopher Simpkins\n" +
		" MIT License\n\n" +
		" Source: https://github.com/source-foundry/uni\n" +
		"=================================================\n\n" +
		" Usage:\n" +
		"  $ uni [glyph 1]...[glyph n]\n\n" +
		" Options:\n" +
		" -h, --help           Application help\n" +
		"     --usage          Application usage\n" +
		" -v, --version        Application version\n\n"
)

func main() {

	// test for at least one argument on command line
	if len(os.Args) < 2 {
		os.Stderr.WriteString("[Error] Please include at least one argument for your Unicode code point search\n")
		os.Stderr.WriteString(USAGE)
		os.Exit(1)
	}

	// define available command line flags
	var versionShort = flag.Bool("v", false, "Application version")
	var versionLong = flag.Bool("version", false, "Application version")
	var helpShort = flag.Bool("h", false, "Help")
	var helpLong = flag.Bool("help", false, "Help")
	var usageLong = flag.Bool("usage", false, "Usage")
	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort:
		os.Stdout.WriteString("uni v" + VERSION + "\n")
		os.Exit(0)
	case *versionLong:
		os.Stdout.WriteString("uni v" + VERSION + "\n")
		os.Exit(0)
	case *helpShort:
		os.Stdout.WriteString(HELP)
		os.Exit(0)
	case *helpLong:
		os.Stdout.WriteString(HELP)
		os.Exit(0)
	case *usageLong:
		os.Stdout.WriteString(USAGE)
		os.Exit(0)
	}

	// display Unicode code point value for glyphs entered as command line arguments
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) > 1 { // handle single argument that includes multiple glyphs
			charList := strings.Split(os.Args[i], "")
			for x := 0; x < len(charList); x++ {
				r, _ := utf8.DecodeRuneInString(charList[x])
				fmt.Printf("%#U\n", r)
			}
		} else { // handle multiple individual arguments
			r, _ := utf8.DecodeRuneInString(os.Args[i])
			fmt.Printf("%#U\n", r)
		}

	}

}
