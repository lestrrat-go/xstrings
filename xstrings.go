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

	b.Grow(len(s))
	b.WriteRune(fn(r))
	b.WriteString(s[n:])
	return b.String()
}

// LcFirst returns a string with the first rune in upper case.
// If for any reason we fail to decode the first rune in the string,
// the same string is returned instead of an error.
func LcFirst(s string) string {
	return convertFirstRune(strings.TrimSpace(s), unicode.ToLower)
}

// UcFirst returns a string with the first rune in upper case.
// If for any reason we fail to decode the first rune in the string,
// the same string is returned instead of an error.
func UcFirst(s string) string {
	return convertFirstRune(strings.TrimSpace(s), unicode.ToUpper)
}

func Snake(s string, options ...SnakeOption) string {
	var delimiter rune = '_'
	var screaming bool
	for _, option := range options {
		switch option.Ident() {
		case identSnakeDelimiter{}:
			delimiter = option.Value().(rune)
		case identSnakeScreaming{}:
			screaming = option.Value().(bool)
		}
	}

	b := getBuilder()
	defer releaseBuilder(b)

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	const (
		isFirst = 1 << (iota + 1) // Only set if this is the first rune
		isBegin
		isLower
		isUpper
		isDigit
		isDelim
	)

	var prev int8 = isFirst
	for len(s) > 0 {
		r, n := utf8.DecodeRuneInString(s)
		s = s[n:]

		var cur int8
		switch {
		case unicode.IsUpper(r):
			cur |= isUpper
			if !screaming {
				r = unicode.ToLower(r)
			}
		case unicode.IsLower(r):
			cur |= isLower
			if screaming {
				r = unicode.ToUpper(r)
			}
		case unicode.IsDigit(r):
			cur |= isDigit
		case unicode.IsSpace(r) || r == delimiter:
			cur |= isDelim
		}

		// special case first letter. it will never be a space, because we
		// already called TrimSpace.
		if prev&isFirst == isFirst {
			cur |= isBegin
			prev = cur
			b.WriteRune(r)
			continue
		}

		// If this is an explcit delimiter (spaces converted to delimiters),
		// we need to make sure that the previous write was not a delimiter.
		// if it was a delimiter, just skip
		if cur&isDelim == isDelim {
			if prev&isDelim == 0 {
				r = delimiter
			}
		} else {
			// If the previous rune was a delimiter, this is going to be the beginning
			if prev&isDelim == isDelim {
				cur |= isBegin
			}

			if cur&isUpper == isUpper && prev&isUpper == isUpper {
				// If the current rune is upper case, and we're in a sequence of uppercase
				// letters, we need to check if the NEXT rune will cause a transition.
				// If a transition is to occur, this current rune belongs in the
				// NEXT word, not the current (e.g. JSONData -> json_data)
				if len(s) > 0 { // nothing to do if we're at the end
					// peek the next rune (TODO: try to reuse this rune in the next iteration?
					r2, _ := utf8.DecodeRuneInString(s)
					if unicode.IsLower(r2) {
						b.WriteRune(delimiter)
						cur |= isBegin
					}
				}
			} else if prev != cur {
				// Insert a delimiter if
				//  * previous state is not the same as current state
				//  * previous r is not a delimiter
				if prev&isDelim == 0 &&
					(prev&isDigit != cur&isDigit ||
						(prev&(isLower|isUpper) != cur&(isLower|isUpper) && /* we transitioned from upper/lower case */
							prev&isBegin == 0 /* ...and the previous write was not the beginning of a word */)) {
					b.WriteRune(delimiter)
					cur |= isBegin
				}
			}
		}
		b.WriteRune(r)
		prev = cur
	}

	return b.String()
}
