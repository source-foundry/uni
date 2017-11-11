# uni

[![Build Status](https://semaphoreci.com/api/v1/sourcefoundry/uni/branches/master/badge.svg)](https://semaphoreci.com/sourcefoundry/uni) [![Go Report Card](https://goreportcard.com/badge/github.com/source-foundry/uni)](https://goreportcard.com/report/github.com/source-foundry/uni)

## About

uni is an application that displays Unicode code points for glyphs included as arguments on the command line.

## Install

uni is developed in Go and compiled to the command line executable `uni`.  A variety of cross-compiled binaries are available for use on Linux, macOS, and Windows systems, or you can download the source and compile the application yourself.  Instructions for both approaches follow.

### Install a pre-compiled binary

[Download the appropriate archive file for your system from the repository releases](https://github.com/source-foundry/uni/releases/latest).  Unpack the `.zip` archive and move the `uni` executable to the desired directory.  For *.nix users (including macOS), the `uni` executable can be placed on your system PATH (e.g. `/usr/local/bin`) to enable use system-wide use with the following:

```
$ uni [args]
```

If you do not install on your system PATH, navigate to the directory where you saved the `uni` executable and use the following:

```
$ ./uni [args]
```

### Compile from source files and install

If you would prefer to build the application from the source, follow these instructions:

- Install [Go](https://golang.org/doc/install)
- [Define your GOPATH](https://github.com/golang/go/wiki/Setting-GOPATH)
- `go get github.com/source-foundry/uni`
- `go install $GOPATH/src/github.com/source-foundry/uni`

## Usage

uni takes glyph arguments and displays the associated Unicode code points.  You can include the glyphs in a single string or separate them with spaces.  Use quotes around special shell characters.

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

## Issues

Please [file an issue report](https://github.com/source-foundry/uni/issues/new) on the repository for any problems that arise with use.

## Contributing

Contributions to the project are encouraged and welcomed.

## License

[MIT License](https://github.com/source-foundry/uni/blob/master/LICENSE)
