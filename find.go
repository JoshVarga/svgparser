package svgparser

// FindChildByID finds the first child with the specified ID.
func (e *Element) FindChildByID(id string) *Element {
	return nil
}

// FindAllChildren finds all children with the given name.
func (e *Element) FindAllChildren(name string) []*Element {
	var elements []*Element
	for _, child := range e.Children {
		if child.Name == name {
			elements = append(elements, child)
		}
		elements = append(elements, child.FindAllChildren(name)...)
	}
	return elements
}
