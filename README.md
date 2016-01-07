# svgparser

Library for parsing and manipulating SVG files.

### Installation

	go get github.com/catiepg/svgparser

### Example

	func ExampleParse() {
		svg := `
			<svg width="100" height="100">
				<circle cx="50" cy="50" r="40" fill="red" />
			</svg>
		`
		reader := strings.NewReader(svg)

		element, _ := svgparser.Parse(reader)

		fmt.Printf("SVG width: %s", element.Attributes["width"])
		fmt.Printf("Circle fill: %s", element.Children[0].Attributes["fill"])

		// Output:
		// SVG width: 100
		// Circle fill: red
	}

### License

The MIT License (MIT)

Copyright (c) 2015 Ekaterina Goranova

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
