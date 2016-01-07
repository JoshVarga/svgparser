package svgparser

import "io"

type Element struct {
	Name       string
	Attributes map[string]string
	Children   []*Element
}

func Parse(source io.Reader) (*Element, error) {
	return nil, nil
}
