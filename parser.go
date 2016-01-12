package svgparser

import (
	"encoding/xml"
	"io"
)

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

func (e *Element) SetAttributes(token xml.StartElement) {
	attributes := make(map[string]string)
	for _, attr := range token.Attr {
		attributes[attr.Name.Local] = attr.Value
	}
	e.Name = token.Name.Local
	e.Attributes = attributes
	e.isPopulated = true
}

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
// SVG root element is not required.
func Parse(source io.Reader) (*Element, error) {
	decoder := xml.NewDecoder(source)
	element := &Element{}
	if err := element.Decode(decoder); err != nil && err != io.EOF {
		return nil, err
	}
	return element, nil
}
