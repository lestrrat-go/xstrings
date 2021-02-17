package xstrings_test

import (
	"testing"

	"github.com/lestrrat-go/xstrings"
	"github.com/stretchr/testify/assert"
)

func TestXstrings(t *testing.T) {
	t.Run("String Functions", func(t *testing.T) {
		testcases := []struct {
			Name     string
			Src      string
			Expected string
			Func     func(string) string
		}{
			{
				Name:     "LcFirst",
				Src:      "Hello, World!",
				Expected: "hello, World!",
				Func: func(s string) string {
					return xstrings.LcFirst(s)
				},
			},
			{
				Name:     "UcFirst",
				Src:      "hello, World!",
				Expected: "Hello, World!",
				Func: func(s string) string {
					return xstrings.UcFirst(s)
				},
			},
		}
		for _, tc := range testcases {
			tc := tc
			t.Run(tc.Name, func(t *testing.T) {
				if !assert.Equal(t, tc.Expected, tc.Func(tc.Src), `values should match`) {
					return
				}
			})
		}
	})
}
