package svgparser_test

import (
	"testing"

	"github.com/catiepg/svgparser"
)

func TestParser(t *testing.T) {
	var testCases = []struct {
		svg     string
		element svgparser.Element
	}{
		{
			`
		<svg width="100" height="100">
			<circle cx="50" cy="50" r="40" fill="red" />
		</svg>
		`,
			svgparser.Element{
				Name: "svg",
				Attributes: map[string]string{
					"width":  "100",
					"height": "100",
				},
				Children: []*svgparser.Element{
					element("circle", map[string]string{"cx": "50", "cy": "50", "r": "40", "fill": "red"}),
				},
			},
		},
		{
			`
		<svg height="400" width="450">
			<g stroke="black" stroke-width="3" fill="black">
				<path id="AB" d="M 100 350 L 150 -300" stroke="red" />
				<path id="BC" d="M 250 50 L 150 300" stroke="red" />
				<path d="M 175 200 L 150 0" stroke="green" />
			</g>
		</svg>
		`,
			svgparser.Element{
				Name: "svg",
				Attributes: map[string]string{
					"width":  "450",
					"height": "400",
				},
				Children: []*svgparser.Element{
					&svgparser.Element{
						Name: "g",
						Attributes: map[string]string{
							"stroke":       "black",
							"stroke-width": "3",
							"fill":         "black",
						},
						Children: []*svgparser.Element{
							element("path", map[string]string{"id": "AB", "d": "M 100 350 L 150 -300", "stroke": "red"}),
							element("path", map[string]string{"id": "BC", "d": "M 250 50 L 150 300", "stroke": "red"}),
							element("path", map[string]string{"d": "M 175 200 L 150 0", "stroke": "green"}),
						},
					},
				},
			},
		},
		{
			"",
			svgparser.Element{},
		},
	}

	for i, test := range testCases {
		actual, err := parse(test.svg, false)

		if !(test.element.Compare(actual) && err == nil) {
			t.Errorf("Parse: expected %v, actual %v\n", i, test.element, actual)
		}
	}
}

func TestValidation(t *testing.T) {
	svg := `
		<circle cx="40" cy="40" r="20">
			<path id="AB" d="M 100 350 L 150 -300" stroke="red" />
		</circle>
		`
	expected := "Document validation error"
	element, err := parse(svg, true)

	if !(element == nil && err.Error() == string(expected)) {
		t.Errorf("Validation: expected %v, actual %v\n", expected, err)
	}
}

func TestValidDocument(t *testing.T) {
	svg := `
		<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="svg-root" width="100%" height="100%" viewBox="0 0 480 360">
			<title id="test-title">color-prop-01-b</title>
			<desc id="test-desc">Test that viewer has the basic capability to process the color property</desc>
			<rect id="test-frame" x="1" y="1" width="478" height="358" fill="none" stroke="#000000"/>
		</svg>
		`

	element, err := parse(svg, true)
	if !(element != nil && err == nil) {
		t.Errorf("Validation: expected %v, actual %v\n", nil, err)
	}
}
