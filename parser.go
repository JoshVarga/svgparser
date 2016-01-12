package svgparser

import (
	"encoding/xml"
	"io"
)

// TODO: custom errors

// Element is a representation of an SVG element.
type Element struct {
	Name        string
	Attributes  map[string]string
	Children    []*Element
	isPopulated bool
}

// FindChildByID finds the first child of the element with the specified ID.
func (e *Element) FindChildByID(id string) *Element {
	return nil
}

// FindAllChildren finds all children of the parent element by element name.
func (e *Element) FindAllChildren(name string) []*Element {
	return []*Element{}
}

// SetAttributes sets the attributes of token to the element.
func (e *Element) SetAttributes(token xml.StartElement) {
	attributes := make(map[string]string)
	for _, attr := range token.Attr {
		attributes[attr.Name.Local] = attr.Value
	}
	e.Name = token.Name.Local
	e.Attributes = attributes
	e.isPopulated = true
}

// Decode adds the attributes of the next start token to the element
// and decodes its child elements.
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
			e.SetAttributes(element)

			nextElement := &Element{}
			err := nextElement.Decode(decoder)
			if err != nil {
				return err
			}

			if nextElement.isPopulated {
				e.Children = append(e.Children, nextElement)
			}

		case xml.CharData:
			// TODO: investigate if any SVG elements can have content
			//       else: validation error
		}
	}

	return nil
}

// Parse creates an Element instance from an SVG input.
func Parse(source io.Reader) (*Element, error) {
	decoder := xml.NewDecoder(source)
	element := &Element{}
	if err := element.Decode(decoder); err != nil && err != io.EOF {
		return nil, err
	}
	return element, nil
}
