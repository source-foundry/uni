// uni is a command line executable that performs Unicode code point to glyph and glyph to Unicode code point searches
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

const (
	version = "2.0.0"

	usage = `Usage: uni (options) [arg 1]...[arg n]
Line Filter Usage: [application command] | uni (options)
`

	help = `=================================================
 uni
 Unicode v10.0 search tool
 Copyright 2018 Christopher Simpkins
 MIT License

 Source: https://github.com/source-foundry/uni
=================================================
 Usage:
  - With command line arguments:
      $ uni (options) [arg 1]...[arg n]
  - As line filter:
      $ [application command] | uni (options)
 Default:
   Search for Unicode code point(s) with glyph(s)
 Options:
  -g, --glyph          Search for glyph(s) with Unicode code point(s)
  -h, --help           Application help
      --usage          Application usage
  -v, --version        Application version
`
)

var versionShort, versionLong, helpShort, helpLong, usageLong, glyphShort, glyphLong *bool
//var listShort, listLong, listCommas, listNewLine *bool
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
	//listShort = flag.Bool("l", false, "List")
	//listLong = flag.Bool("list", false, "List")
	//listCommas = flag.Bool("comma", false, "Comma delimiters in lists")
	//listNewLine = flag.Bool("newline", false, "Newline delimiters in lists")
}

func main() {

	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		fmt.Printf("uni v%s\n", version)
		fmt.Printf("Unicode Standard v%s\n", unicode.Version)
		os.Exit(0)
	case *helpShort, *helpLong:
		fmt.Println(help)
		os.Exit(0)
	case *usageLong:
		fmt.Println(usage)
		os.Exit(0)
	}

	//if *listShort || *listLong {
	//	if len(os.Args) < 3 {
	//		os.Stderr.WriteString("[Error] Please include at least one Unicode code point range argument for the list.\n")
	//		os.Exit(1)
	//	}
	//	var finalUniList []string
	//	for _, x := range flag.Args() {
	//		if strings.Contains(x, "-") {
	//			tempUniList := strings.Split(x, "-")
	//			startInt, startErr := strconv.ParseInt(tempUniList[0], 16, 32)
	//			endInt, endErr := strconv.ParseInt(tempUniList[1], 16, 32)
	//			if startErr != nil {
	//				// TODO
	//			}
	//			if endErr != nil {
	//				// TODO
	//			}
	//			for i := startInt; i <= endInt; i++ {
	//				finalUniList = append(finalUniList, fmt.Sprintf("%U", i))
	//
	//			}
	//		}
	//	}
	//
	//	if *listCommas {
	//		fmt.Print(strings.Join(finalUniList, ","))
	//	} else if *listNewLine {
	//		fmt.Print(strings.Join(finalUniList, "\n"))
	//	} else {
	//		fmt.Print(strings.Join(finalUniList, " "))
	//	}
	//
	//	// TODO: need to exit to prevent code below from executing
	//	os.Exit(0)
	//}

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
