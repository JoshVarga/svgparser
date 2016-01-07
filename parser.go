package svgparser

import "io"

type Element struct {
	Name       string
	Attributes map[string]string
	Children   []*Element
}

func (e *Element) FindChildByID(id string) *Element {
	return nil
}

func (e *Element) FindAllChildren(name string) []*Element {
	return []*Element{}
}

func Parse(source io.Reader) (*Element, error) {
	return nil, nil
}
