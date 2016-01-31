package svgparser

import (
	"encoding/xml"
	"io"
)

// TODO: custom errors

// Element is a representation of an SVG element.
type Element struct {
	Name       string
	Attributes map[string]string
	Children   []*Element
}

// NewElement creates element from decoder token.
func NewElement(token xml.StartElement) *Element {
	element := &Element{}
	attributes := make(map[string]string)
	for _, attr := range token.Attr {
		attributes[attr.Name.Local] = attr.Value
	}
	element.Name = token.Name.Local
	element.Attributes = attributes
	return element
}

// FindChildByID finds the first child of the element with the specified ID.
func (e *Element) FindChildByID(id string) *Element {
	return nil
}

// FindAllChildren finds all children of the parent element by element name.
func (e *Element) FindAllChildren(name string) []*Element {
	return []*Element{}
}

// Compare compares two elements.
func (e *Element) Compare(o *Element) bool {
	if e.Name != o.Name ||
		len(e.Attributes) != len(o.Attributes) ||
		len(e.Children) != len(o.Children) {
		return false
	}

	for k, v := range e.Attributes {
		if v != o.Attributes[k] {
			return false
		}
	}

	for i, child := range e.Children {
		if !child.Compare(o.Children[i]) {
			return false
		}
	}
	return true
}

// DecodeFirst creates the first element from the decoder.
func DecodeFirst(decoder *xml.Decoder) (*Element, error) {
	for {
		token, err := decoder.Token()
		if token == nil && err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		switch element := token.(type) {
		case xml.StartElement:
			return NewElement(element), nil
		}
	}
	// TODO: no start element
	return nil, nil
}

// Decode decodes the child elements of element.
func (e *Element) Decode(decoder *xml.Decoder) error {
	for {
		token, err := decoder.Token()
		if token == nil && err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch element := token.(type) {
		case xml.StartElement:
			nextElement := NewElement(element)
			err := nextElement.Decode(decoder)
			if err != nil {
				return err
			}

			e.Children = append(e.Children, nextElement)

		case xml.CharData:
			// TODO: investigate if any SVG elements can have content
			//       else: validation error

		case xml.EndElement:
			if element.Name.Local == e.Name {
				return nil
			}
		}
	}

	return nil
}

// Parse creates an Element instance from an SVG input.
func Parse(source io.Reader) (*Element, error) {
	decoder := xml.NewDecoder(source)
	element, err := DecodeFirst(decoder)
	if err != nil {
		return nil, err
	}
	if err := element.Decode(decoder); err != nil && err != io.EOF {
		return nil, err
	}
	return element, nil
}
