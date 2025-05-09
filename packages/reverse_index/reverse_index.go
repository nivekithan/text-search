package reverse_index

import (
	"github.com/nivekithan/text-search/packages/common"
	"github.com/nivekithan/text-search/packages/tokeniser"
)

type ReverseIndex struct {
	tokeniser tokeniser.Tokeniser
	index     map[tokeniser.Token]common.Set[int]
}

func NewReverseIndex(langTokeniser tokeniser.Tokeniser) *ReverseIndex {
	return &ReverseIndex{
		tokeniser: langTokeniser,
		index:     map[tokeniser.Token]common.Set[int]{},
	}
}

func (r *ReverseIndex) AddDocument(document string, id int) {
	tokens := r.tokeniser.Tokens(document)

	for _, token := range tokens {
		existingSet, ok := r.index[token]

		if !ok {
			newSet := common.NewSet[int]()

			newSet.Add(id)
			r.index[token] = *newSet
			continue
		}

		existingSet.Add(id)

	}
}

func (r *ReverseIndex) SearchToken(token tokeniser.Token) []int {
	existingsIds, ok := r.index[token]

	if !ok {
		return []int{}
	}

	return existingsIds.Values()
}
