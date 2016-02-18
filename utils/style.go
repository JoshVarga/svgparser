package utils

import "regexp"

type Style struct {
	Property string
	Value    string
}

type Styles []*Style

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

func StyleParser(raw string) Styles {
	var styles Styles
	regex := regexp.MustCompile(`([a-zA-z-]+):([^:;]+);?`)
	for _, s := range regex.FindAllStringSubmatch(raw, -1) {
		styles = append(styles, &Style{s[1], s[2]})
	}
	return styles
}
