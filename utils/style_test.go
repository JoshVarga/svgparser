package utils_test

import (
	"testing"

	"."
)

func TestStyleParser(t *testing.T) {
	var testCases = []struct {
		styles   string
		expected utils.Styles
	}{
		{
			"fill:white;stroke:#000000;",
			utils.Styles(
				[]*utils.Style{
					&utils.Style{"fill", "white"},
					&utils.Style{"stroke", "#000000"},
				},
			),
		},
		{
			"fill:white;stroke-opacity:1",
			utils.Styles(
				[]*utils.Style{
					&utils.Style{"fill", "white"},
					&utils.Style{"stroke-opacity", "1"},
				},
			),
		},
	}

	for _, test := range testCases {
		styles := utils.StyleParser(test.styles)
		if !test.expected.Compare(styles) {
			t.Errorf("Style: expected %v, actual %v\n", test.expected, styles)
		}
	}
}
