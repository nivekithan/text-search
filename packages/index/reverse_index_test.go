package index

import (
	"testing"

	"github.com/nivekithan/text-search/packages/tokeniser"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	tokeniser := tokeniser.NewEnglishTokeniser()
	reverseIndex := NewIndex(tokeniser)

	testCases := map[string]struct {
		documents []string
		queries   []struct {
			query    string
			expected []int
		}
	}{
		"simple search": {
			documents: []string{
				"I am nivekithan",               // 0
				"Hey there",                     // 1
				"How are you",                   // 2
				"Where are you from ",           // 3
				"from which place are you from", // 4
				"which food you like",           // 5
				"hello world",                   // 6
			},
			queries: []struct {
				query    string
				expected []int
			}{
				{query: "hello", expected: []int{6}},
				{query: "world", expected: []int{6}},
				{query: "nivekithan", expected: []int{0}},
				{query: "from", expected: []int{3, 4}},
				{query: "which", expected: []int{4, 5}},
				{query: "are", expected: []int{2, 3, 4}},
				{query: "random", expected: []int{}},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			for id, document := range testCase.documents {
				reverseIndex.AddDocument(document, id)
			}

			for _, query := range testCase.queries {
				t.Run(query.query, func(t *testing.T) {
					tokens := tokeniser.Tokens(query.query)

					if len(tokens) != 1 {
						t.Errorf("expected 1 token, got %v", len(tokens))
					}

					token := tokens[0]

					actual := reverseIndex.SearchToken(token)

					assert.ElementsMatch(t, actual, query.expected)
				})
			}
		})
	}
}
