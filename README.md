# uni

[![Build Status](https://semaphoreci.com/api/v1/sourcefoundry/uni/branches/master/badge.svg)](https://semaphoreci.com/sourcefoundry/uni) [![Build status](https://ci.appveyor.com/api/projects/status/1cbesa0qnbqgoj0p/branch/master?svg=true)](https://ci.appveyor.com/project/chrissimpkins/uni/branch/master) [![Go Report Card](https://goreportcard.com/badge/github.com/source-foundry/uni)](https://goreportcard.com/report/github.com/source-foundry/uni)

## About

`uni` is an application that displays Unicode code points for glyphs that are either included as arguments on the command line or piped through the standard input stream to the `uni` executable.

<img src="https://raw.githubusercontent.com/source-foundry/uni/img/img/uni-crunch.png" alt="uni Argument Example" width="99%"/>

<img src="https://raw.githubusercontent.com/source-foundry/uni/img/img/uni2-crunch.png" alt="uni Line Filter Example" width="99%"/>

## Install

uni is developed in Go and compiled to the command line executable `uni` (`uni.exe` on Windows).  A variety of cross-compiled binaries are available for use on Linux, macOS, and Windows systems, or you can download the source and compile the application yourself.  Instructions for both approaches follow.

## Installation

### Approach 1: Install the pre-compiled binary executable file

Download the latest compiled release file for your operating system and architecture from [the Releases page](https://github.com/source-foundry/uni/releases/latest).

#### Linux / macOS

Unpack the tar.gz archive and move the `uni` executable file to a directory on your system PATH (e.g. `/usr/local/bin`).  This can be performed by executing the following command in the root of the unpacked archive:

```
$ mv uni /usr/local/bin/uni
```

There are no dependencies contained in the archive.  You can delete all downloaded archive files after the above step.

#### Windows

Unpack the zip archive and move the `uni.exe` executable file to a directory on your system PATH. See [details here](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows) for more information about how to do this.

There are no dependencies contained in the archive.  You can delete all downloaded archive files after the above step.

### Approach 2: Compile from the source code and install

You must install the Go programming language (which includes the `go` tool) in order to compile the project from source.  Follow the [instructions on the Go download page](https://golang.org/dl/) for your platform. 

Once you have installed Go and configured your settings so that Go executables are installed on your system PATH, use the following command to (1) pull the master branch of the ink repository; (2) compile the ink executable from source for your platform/architecture configuration; (3) install the executable on your system:

```
$ go get github.com/source-foundry/uni
```

## Uninstall

The installation includes a single executable binary file.  If you installed with `go get` or added one of the pre-compiled binaries on your system `$PATH` on *.nix systems, you can uninstall with:

```
$ rm $(which uni)
```

## Usage

### Glyphs as arguments to `uni`

`uni` takes glyph arguments and displays the associated Unicode code points.  You can include the glyphs in a single string or separate them with spaces.  Use quotes around special shell characters.

```
$ uni [glyph 1]...[glyph n]
```

#### Example

```
$ uni Aa1Ø€βф▀र༩↵√ナ
U+0041 'A'
U+0061 'a'
U+0031 '1'
U+00D8 'Ø'
U+20AC '€'
U+03B2 'β'
U+0444 'ф'
U+2580 '▀'
U+0930 'र'
U+0F29 '༩'
U+21B5 '↵'
U+221A '√'
U+30CA 'ナ'
```

### Glyphs piped through standard input stream

You can also pipe text data to `uni` through the standard input stream. `uni` will process every glyph that it receives in the stdin stream and print the associated Unicode code point to standard output.

```
$ echo -n "Aa1Ø€βф▀र༩↵√ナ" | uni
U+0041 'A'
U+0061 'a'
U+0031 '1'
U+00D8 'Ø'
U+20AC '€'
U+03B2 'β'
U+0444 'ф'
U+2580 '▀'
U+0930 'र'
U+0F29 '༩'
U+21B5 '↵'
U+221A '√'
U+30CA 'ナ'
```

## Issues

Please [file an issue report](https://github.com/source-foundry/uni/issues/new) on the repository for any problems that arise with use.

## Contributing

Contributions to the project are encouraged and welcomed.

## License

[MIT License](https://github.com/source-foundry/uni/blob/master/LICENSE)
