package xstrings

import "github.com/lestrrat-go/option"

type Option = option.Interface

type identSnakeDelimiter struct{}
type identSnakeScreaming struct{}
type identCamelLower struct{}

type SnakeOption interface {
	Option
	snakeOption()
}

type snakeOption struct{ Option }
func (*snakeOption) snakeOption() {}

type CamelOption interface {
	Option
	camelOption()
}

type camelOption struct{ Option }
func (*camelOption) camelOption() {}

// WithDelimiter allows you to change the delimiter used in snake case coversion
func WithDelimiter(r rune) SnakeOption {
	return &snakeOption{option.New(identSnakeDelimiter{}, r)}
}

// WithScreaming specifies that the conversion to snake case should produce
// words in all uppercase (e.g. SNAKE_CASE)
func WithScreaming(v bool) SnakeOption {
	return &snakeOption{option.New(identSnakeScreaming{}, v)}
}

// WithLowerCamel specifies that the first letter of a camel cased string should
// be lower cased
func WithLowerCamel(v bool) CamelOption {
	return &camelOption{option.New(identCamelLower{}, v)}
}
