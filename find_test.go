package svgparser_test

import (
	"testing"

	"github.com/catiepg/svgparser"
)

func testElement() *svgparser.Element {
	svg := `
		<svg width="1000" height="600">
			<g id="first">
				<rect width="5" height="3" id="inFirst"/>
				<rect width="5" height="2" id="inFirst"/>
			</g>
			<g id="second">
				<path d="M50 50 Q50 100 100 100"/>
				<rect width="5" height="1"/>
			</g>
		</svg>
	`
	element, _ := parse(svg, false)
	return element
}

func equals(t *testing.T, name string, expected, actual *svgparser.Element) {
	if !(expected == actual || expected.Compare(actual)) {
		t.Errorf("%s: expected %v, actual %v\n", name, expected, actual)
	}
}

func equalSlices(t *testing.T, name string, expected, actual []*svgparser.Element) {
	if len(expected) != len(actual) {
		t.Errorf("%s: expected %v, actual %v\n", name, expected, actual)
		return
	}

	for i, r := range actual {
		equals(t, name, expected[i], r)
	}
}

func TestFindAllChildren(t *testing.T) {
	svgElement := testElement()

	equalSlices(t, "Find", []*svgparser.Element{
		element("rect", map[string]string{"width": "5", "height": "3", "id": "inFirst"}),
		element("rect", map[string]string{"width": "5", "height": "2", "id": "inFirst"}),
		element("rect", map[string]string{"width": "5", "height": "1"}),
	}, svgElement.FindAllChildren("rect"))

	equalSlices(t, "Find", []*svgparser.Element{}, svgElement.FindAllChildren("circle"))
}

func TestFindByID(t *testing.T) {
	svgElement := testElement()

	equals(t, "Find", &svgparser.Element{
		Name:       "g",
		Attributes: map[string]string{"id": "second"},
		Children: []*svgparser.Element{
			element("path", map[string]string{"d": "M50 50 Q50 100 100 100"}),
			element("rect", map[string]string{"width": "5", "height": "1"}),
		},
	}, svgElement.FindByID("second"))

	equals(t, "Find",
		element("rect", map[string]string{"width": "5", "height": "3", "id": "inFirst"}),
		svgElement.FindByID("inFirst"),
	)

	equals(t, "Find", nil, svgElement.FindByID("missing"))
}
