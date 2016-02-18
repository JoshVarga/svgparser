package utils

import "regexp"

// Style represents a CSS style property and its value.
type Style struct {
	Property string
	Value    string
}

// Styles is a collection of Style objects.
type Styles []*Style

// Compare compares two Style objects.
func (ss Styles) Compare(o Styles) bool {
	if len(ss) != len(o) {
		return false
	}

	for i, s := range ss {
		if !(s.Property == o[i].Property && s.Value == o[i].Value) {
			return false
		}
	}
	return true
}

// StyleParser takes value of a style attribute and converts it to
// Style objects.
func StyleParser(raw string) Styles {
	var styles Styles
	regex := regexp.MustCompile(`([a-zA-z-]+):([^:;]+);?`)
	for _, s := range regex.FindAllStringSubmatch(raw, -1) {
		styles = append(styles, &Style{s[1], s[2]})
	}
	return styles
}
