package svgparser

// FindID finds the first child with the specified ID.
func (e *Element) FindID(id string) *Element {
	for _, child := range e.Children {
		if childID, ok := child.Attributes["id"]; ok && childID == id {
			return child
		}
		if element := child.FindID(id); element != nil {
			return element
		}
	}
	return nil
}

// FindAll finds all children with the given name.
func (e *Element) FindAll(name string) []*Element {
	var elements []*Element
	for _, child := range e.Children {
		if child.Name == name {
			elements = append(elements, child)
		}
		elements = append(elements, child.FindAll(name)...)
	}
	return elements
}
