package index

import (
	"github.com/nivekithan/text-search/packages/common"
	"github.com/nivekithan/text-search/packages/tokeniser"
)

type Index struct {
	tokeniser     tokeniser.Tokeniser
	reverse_index *common.ReverseIndex[tokeniser.Token]
	trie          *common.Trie[rune]
}

func NewIndex(langTokeniser tokeniser.Tokeniser) *Index {
	return &Index{
		tokeniser:     langTokeniser,
		reverse_index: common.NewReverseIndex[tokeniser.Token](),
		trie:          common.NewTrie('0'),
	}
}

func (r *Index) AddDocument(document string, id int) {
	tokens := r.tokeniser.Tokens(document)

	for _, token := range tokens {
		r.reverse_index.AddEntry(token, id)

		runes := make([]rune, len(token))

		for _, ch := range token {
			runes = append(runes, ch)
		}

		r.trie.Add(runes)

	}
}

func (r *Index) SearchToken(token tokeniser.Token) []int {
	return r.reverse_index.GetEntry(token)
}
