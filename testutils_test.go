package svgparser_test

import (
	"strings"
	"testing"

	"github.com/catiepg/svgparser"
)

func element(name string, attrs map[string]string) *svgparser.Element {
	return &svgparser.Element{
		Name:       name,
		Attributes: attrs,
		Children:   []*svgparser.Element{},
	}
}

func parse(svg string) (*svgparser.Element, error) {
	element, err := svgparser.Parse(strings.NewReader(svg))
	return element, err
}

func equals(t *testing.T, name string, expected, actual *svgparser.Element) {
	if !(expected == actual || expected.Compare(actual)) {
		t.Errorf("%s: expected %v, actual %v\n", name, expected, actual)
	}
}

func equalSlices(t *testing.T, name string, expected, actual []*svgparser.Element) {
	if len(expected) != len(actual) {
		t.Errorf("%s: expected %v, actual %v\n", name, expected, actual)
		return
	}

	for i, r := range actual {
		equals(t, name, expected[i], r)
	}
}
