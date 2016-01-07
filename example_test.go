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
