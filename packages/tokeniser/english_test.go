package tokeniser

import (
	"reflect"
	"testing"
)

func TestEnglishTokeniser(t *testing.T) {

	testCases := map[string][]string{
		"Hey there":                           {"hey", "there"},
		"Hello, World!":                       {"hello,", "world!"},
		"Lot                      off spaces": {"lot", "off", "spaces"},
	}

	for input, expected := range testCases {
		tokeniser := &EnglishTokeniser{}
		actual := tokeniser.Tokens(input)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}
