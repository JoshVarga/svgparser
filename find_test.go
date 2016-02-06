package svgparser_test

import (
	"strings"
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
	reader := strings.NewReader(svg)
	element, _ := svgparser.Parse(reader)
	return element
}

func TestFindAllChildren(t *testing.T) {
	svgElement := testElement()

	equalSlices(t, []*svgparser.Element{
		element("rect", map[string]string{"width": "5", "height": "3", "id": "inFirst"}),
		element("rect", map[string]string{"width": "5", "height": "2", "id": "inFirst"}),
		element("rect", map[string]string{"width": "5", "height": "1"}),
	}, svgElement.FindAllChildren("rect"))

	equalSlices(t, []*svgparser.Element{}, svgElement.FindAllChildren("circle"))
}

func TestFindByID(t *testing.T) {
	svgElement := testElement()

	equals(t, &svgparser.Element{
		Name:       "g",
		Attributes: map[string]string{"id": "second"},
		Children: []*svgparser.Element{
			element("path", map[string]string{"d": "M50 50 Q50 100 100 100"}),
			element("rect", map[string]string{"width": "5", "height": "1"}),
		},
	}, svgElement.FindByID("second"))

	equals(t,
		element("rect", map[string]string{"width": "5", "height": "3", "id": "inFirst"}),
		svgElement.FindByID("inFirst"),
	)

	equals(t, nil, svgElement.FindByID("missing"))
}
