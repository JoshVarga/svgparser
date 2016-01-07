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

	fmt.Println(element.Attributes["width"])
	fmt.Println(element.Children[0].Attributes["fill"])

	// Output:
	// 100
	// red
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

	fmt.Println(rectangles[0].Attributes["fill"])
	fmt.Println(rectangles[1].Attributes["y"])

	// Output:
	// #000
	// 1
}

func ExampleElement_FindChildByID() {
	svg := `
		<svg width="1000" height="600">
			<rect id="black" fill="#000" width="5" height="3"/>
			<rect id="gray" fill="#888" width="5" height="2" y="1"/>
			<rect id="white" fill="#fff" width="5" height="1" y="2"/>
		</svg>
	`
	reader := strings.NewReader(svg)
	element, _ := svgparser.Parse(reader)

	white := element.FindChildByID("white")

	fmt.Println(white.Attributes["fill"])

	// Output:
	// #fff
}
