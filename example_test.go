package svgparser_test

import (
	"fmt"
	"github.com/catiepg/svgparser"
	"strings"
)

func ExampleParse() {
	svg := `
		<svg width="100" height="100">
			<circle cx="50" cy="50" r="40" fill="red" />
		</svg>
	`
	reader := strings.NewReader(svg)

	element, _ := svgparser.Parse(reader)

	fmt.Printf("SVG width: %s\n", element.Attributes["width"])
	fmt.Printf("Circle fill: %s", element.Children[0].Attributes["fill"])

	// Output:
	// SVG width: 100
	// Circle fill: red
}

func ExampleElement_FindAllChildren() {
	svg := `
		<svg width="1000" height="600">
			<rect fill="#000" width="5" height="3"/>
			<rect fill="#444" width="5" height="2" y="1"/>
			<rect fill="#888" width="5" height="1" y="2"/>
		</svg>
	`
	reader := strings.NewReader(svg)
	element, _ := svgparser.Parse(reader)

	rectangles := element.FindAllChildren("rect")

	fmt.Printf("First child fill: %s\n", rectangles[0].Attributes["fill"])
	fmt.Printf("Second child height: %s", rectangles[1].Attributes["height"])

	// Output:
	// First child fill: #000
	// Second child height: 2
}

func ExampleElement_FindByID() {
	svg := `
		<svg width="1000" height="600">
			<rect id="black" fill="#000" width="5" height="3"/>
			<rect id="gray" fill="#888" width="5" height="2" y="1"/>
			<rect id="white" fill="#fff" width="5" height="1" y="2"/>
		</svg>
	`
	reader := strings.NewReader(svg)
	element, _ := svgparser.Parse(reader)

	white := element.FindByID("white")

	fmt.Printf("White rect fill: %s", white.Attributes["fill"])

	// Output:
	// White rect fill: #fff
}
