package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie('0')

	testcases := map[string]struct {
		terms  []string
		prefix string
		result []string
	}{
		"simple search": {
			terms:  []string{"hello", "hell", "heaven", "random"},
			prefix: "he",
			result: []string{"hello", "hell", "heaven"},
		},
		"empty prefix": {
			terms:  []string{"hello", "hell", "heaven"},
			prefix: "",
			result: []string{},
		},
		"exact match": {
			terms:  []string{"hello", "hell", "heaven"},
			prefix: "heaven",
			result: []string{"heaven"},
		},
		"prefix matches nothing": {
			terms:  []string{"hello", "hell", "heaven"},
			prefix: "random",
			result: []string{},
		},
		"prefix matches extact and partial": {
			terms:  []string{"hello", "hell", "heaven"},
			prefix: "hell",
			result: []string{"hell", "hello"},
		},
	}

	for name, testcase := range testcases {
		t.Run(name, func(t *testing.T) {
			for _, term := range testcase.terms {
				trie.Add([]rune(term))
			}

			results, err := trie.Search([]rune(testcase.prefix))
			if err != nil {
				t.Errorf("error searching for prefix: %v", err)
			}

			resultsInStrings := []string{}

			for _, result := range results {
				resultsInStrings = append(resultsInStrings, string(result))
			}

			assert.ElementsMatch(t, resultsInStrings, testcase.result)
		})
	}
}
