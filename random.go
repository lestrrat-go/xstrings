package xstrings

import (
	"math/rand"
	"strings"
)

type RandomStringBuilder struct {
	elements []rune
}

func NewRandomStringBuilder(elements ...rune) *RandomStringBuilder {
	return &RandomStringBuilder{
		elements: elements,
	}
}

func (rs *RandomStringBuilder) Build(size int) string {
	elems := rs.elements
	rslen := len(elems)

	var sb strings.Builder
	for i := 0; i < size; i++ {
		sb.WriteRune(elems[rand.Intn(rslen)])
	}
	return sb.String()
}
