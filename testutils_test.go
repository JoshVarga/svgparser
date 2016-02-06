package svgparser_test

import (
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

func equals(t *testing.T, expected, actual *svgparser.Element) {
	if !(expected == actual || expected.Compare(actual)) {
		t.Errorf("Find: expected %v, actual %v\n", expected, actual)
	}
}

func equalSlices(t *testing.T, expected, actual []*svgparser.Element) {
	if len(expected) != len(actual) {
		t.Errorf("Find: expected %v, actual %v\n", expected, actual)
		return
	}

	for i, r := range actual {
		equals(t, expected[i], r)
	}
}
