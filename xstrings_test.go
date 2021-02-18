package xstrings_test

import (
	"strings"
	"testing"

	"github.com/lestrrat-go/xstrings"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	Src      string
	Expected string
}

func expectInLowerCamel(cases []Case) []Case {
	ret := make([]Case, len(cases))
	for i, c := range cases {
		ret[i].Src = c.Src

		switch c.Expected {
		case "JSONData":
			ret[i].Expected = "jsonData"
		case "ID":
			ret[i].Expected = "id"
		default:
			ret[i].Expected = xstrings.LcFirst(c.Expected)
		}
	}
	return ret
}

func expectInAllUpper(cases []Case) []Case {
	ret := make([]Case, len(cases))
	for i, c := range cases {
		ret[i].Src = c.Src
		ret[i].Expected = strings.ToUpper(c.Expected)
	}
	return ret
}

func expectInAlternateDelimiter(cases []Case, delim string) []Case {
	// In this case we need to change the src's occurrences as well
	ret := make([]Case, len(cases))
	for i, c := range cases {
		ret[i].Src = strings.ReplaceAll(c.Src, "_", delim)
		ret[i].Expected = strings.ReplaceAll(c.Expected, "_", delim)
	}
	return ret
}

func TestXstrings(t *testing.T) {
	t.Run("String Functions", func(t *testing.T) {
		camelTestCases := []Case{
			{"test_case", "TestCase"},
			{"test.case", "TestCase"},
			{"test", "Test"},
			{"TestCase", "TestCase"},
			{" test  case ", "TestCase"},
			{"", ""},
			{"many_many_words", "ManyManyWords"},
			{"AnyKind of_string", "AnyKindOfString"},
			{"odd-fix", "OddFix"},
			{"numbers2And55with000", "Numbers2And55With000"},
			{"ID", "ID"},
			{"json_data", "JSONData"},
		}
		snakeTestCases := []Case{
			{"testCase", "test_case"},
			{"TestCase", "test_case"},
			{"Test Case", "test_case"},
			{" Test Case", "test_case"},
			{"Test Case ", "test_case"},
			{" Test Case ", "test_case"},
			{"test", "test"},
			{"test_case", "test_case"},
			{"Test", "test"},
			{"", ""},
			{"ManyManyWords", "many_many_words"},
			{"manyManyWords", "many_many_words"},
			{"AnyKind of_string", "any_kind_of_string"},
			{"numbers2and55with000", "numbers_2_and_55_with_000"},
			{"JSONData", "json_data"},
			{"userID", "user_id"},
			{"AAAbbb", "aa_abbb"},
			{"1A2", "1_a_2"},
			{"A1B", "a_1_b"},
			{"A1A2A3", "a_1_a_2_a_3"},
			{"A1 A2 A3", "a_1_a_2_a_3"},
			{"AB1AB2AB3", "ab_1_ab_2_ab_3"},
			{"AB1 AB2 AB3", "ab_1_ab_2_ab_3"},
			{"some string", "some_string"},
			{" some string", "some_string"},
		}
		testcases := []struct {
			Name  string
			Func  func(string) string
			Cases []Case
		}{
			{
				Name: "LcFirst",
				Func: xstrings.LcFirst,
				Cases: []Case{
					{
						Src:      "Hello, World!",
						Expected: "hello, World!",
					},
				},
			},
			{
				Name: "UcFirst",
				Func: xstrings.UcFirst,
				Cases: []Case{
					{
						Src:      "hello, World!",
						Expected: "Hello, World!",
					},
				},
			},
			{
				Name:  "Snake",
				Func:  func(s string) string { return xstrings.Snake(s) },
				Cases: snakeTestCases,
			},
			{
				Name:  "Screaming Snake",
				Func:  func(s string) string { return xstrings.Snake(s, xstrings.WithScreaming(true)) },
				Cases: expectInAllUpper(snakeTestCases),
			},
			{
				Name:  "Screaming Snake",
				Func:  func(s string) string { return xstrings.Snake(s, xstrings.WithDelimiter('~')) },
				Cases: expectInAlternateDelimiter(snakeTestCases, "~"),
			},
			{
				Name:  "Camel",
				Func:  func(s string) string { return xstrings.Camel(s) },
				Cases: camelTestCases,
			},
			{
				Name:  "Camel Lower",
				Func:  func(s string) string { return xstrings.Camel(s, xstrings.WithLowerCamel(true)) },
				Cases: expectInLowerCamel(camelTestCases),
			},
		}
		for _, tc := range testcases {
			tc := tc
			t.Run(tc.Name, func(t *testing.T) {
				for _, c := range tc.Cases {
					c := c
					t.Run(c.Src, func(t *testing.T) {
						if !assert.Equal(t, c.Expected, tc.Func(c.Src), `values should match`) {
							return
						}
					})
				}
			})
		}
	})
}
