package xstrings

import "github.com/lestrrat-go/option"

type Option = option.Interface

type identSnakeDelimiter struct{}
type identSnakeScreaming struct{}

type SnakeOption interface {
	Option
	snakeOption()
}

type snakeOption struct{ Option }

func (*snakeOption) snakeOption() {}

// WithDelimiter allows you to change the delimiter used in snake case coversion
func WithDelimiter(r rune) SnakeOption {
	return &snakeOption{option.New(identSnakeDelimiter{}, r)}
}

// WithScreaming specifies that the conversion to snake case should produce
// words in all uppercase (e.g. SNAKE_CASE)
func WithScreaming(v bool) SnakeOption {
	return &snakeOption{option.New(identSnakeScreaming{}, v)}
}
