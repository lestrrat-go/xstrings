package xstrings

import (
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

var builderPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func getBuilder() *strings.Builder {
	return builderPool.Get().(*strings.Builder)
}

func releaseBuilder(v *strings.Builder) {
	v.Reset()
	builderPool.Put(v)
}

func convertFirstRune(s string, fn func(rune) rune) string {
	if utf8.RuneCountInString(s) == 0 {
		return s
	}

	r, n := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return s
	}

	b := getBuilder()
	defer releaseBuilder(b)
	b.WriteRune(fn(r))
	b.WriteString(s[n:])
	return b.String()
}

// LcFirst returns a string with the first rune in upper case.
// If for any reason we fail to decode the first rune in the string,
// the same string is returned instead of an error.
func LcFirst(s string) string {
	return convertFirstRune(s, unicode.ToLower)
}

// UcFirst returns a string with the first rune in upper case.
// If for any reason we fail to decode the first rune in the string,
// the same string is returned instead of an error.
func UcFirst(s string) string {
	return convertFirstRune(s, unicode.ToUpper)
}
