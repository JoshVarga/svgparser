package svgparser_test

import (
	"strings"
	"testing"

	"github.com/catiepg/svgparser"
)

func element(name string, attrs map[string]string) *svgparser.Element {
	return &svgparser.Element{
		Name:       name,
		Attributes: attrs,
		Children:   []*svgparser.Element{},
	}
}

func compareSlices(t *testing.T, expected, actual []*svgparser.Element) {
	if len(expected) != len(actual) {
		t.Errorf("Find: expected %v, actual %v\n", expected, actual)
		return
	}

	for i, r := range actual {
		if !expected[i].Compare(r) {
			t.Errorf("Find: expected %v, actual %v\n", expected[i], r)
		}
	}
}

func TestFindAllChildren(t *testing.T) {
	svg := `
		<svg width="1000" height="600">
			<g>
				<rect width="5" height="3"/>
				<rect width="5" height="2"/>
			</g>
			<g>
				<circle r="10" cx="20" cy="30"/>
				<rect width="5" height="1"/>
			</g>
		</svg>
	`
	reader := strings.NewReader(svg)
	svgElement, _ := svgparser.Parse(reader)

	compareSlices(t, []*svgparser.Element{
		element("rect", map[string]string{"width": "5", "height": "3"}),
		element("rect", map[string]string{"width": "5", "height": "2"}),
		element("rect", map[string]string{"width": "5", "height": "1"}),
	}, svgElement.FindAllChildren("rect"))

	compareSlices(t, []*svgparser.Element{}, svgElement.FindAllChildren("path"))
}
