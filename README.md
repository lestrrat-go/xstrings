# github.com/lestrrat-go/xstrings ![](https://github.com/lestrrat-go/jwx/workflows/CI/badge.svg) [![Go reference](https://pkg.go.dev/badge/github.com/lestrrat-go/xstrings.svg)](https://pkg.go.dev/github.com/lestrrat-go/xstrings) [![codecov.io](http://codecov.io/github/lestrrat-go/xstrings/coverage.svg?branch=main)](http://codecov.io/github/lestrrat-go/xstrings?branch=main)

Unicode-aware string utilities for Go

# DESCRIPTION

Many string utilities available on the web only work with characters in ASCII range.
This library does The Right Thing and goes through the proper `unicode` / `unicode/utf8`
decoding routines before processing the strings

# AVAILABLE FUNCTIONS

| Name | Description |
|------|-------------|
| LcFirst(string) string | Converts first rune of string to lower case |
| UcFirst(string) string | Converts first rune of string to upper case |
| Snake(string, SnakeOptions...) | Converts a string into snake_case. Supports alternate delimiters and SCREAMING case via options |
| Camel(string, CamelOptions...) | Converts a string into CamelCase. Supports lowerCaseCame via options|
| FirstNRunes(strings, int) | Returns a strings consisting of the first N runes in the string|

