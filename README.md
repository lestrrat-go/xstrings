# xstrings

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

