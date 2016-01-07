package svgparser

import "io"

type Element struct {
	Name       string
	Attributes map[string]string
	Children   []*Element
}

// FindChildByID finds the first child of the element with the specified ID.
func (e *Element) FindChildByID(id string) *Element {
	return nil
}

// FindAllChildren finds all children of the parent element by element name.
func (e *Element) FindAllChildren(name string) []*Element {
	return []*Element{}
}

// Parse creates an Element instance from an SVG input.
// SVG root element is not required.
func Parse(source io.Reader) (*Element, error) {
	return nil, nil
}
