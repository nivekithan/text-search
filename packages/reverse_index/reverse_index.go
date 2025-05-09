package reverse_index

import "github.com/nivekithan/text-search/packages/tokeniser"

type ReverseIndex struct {
	tokeniser tokeniser.Tokeniser
	index     map[tokeniser.Token][]int
}

func NewReverseIndex(langTokeniser tokeniser.Tokeniser) *ReverseIndex {
	return &ReverseIndex{
		tokeniser: langTokeniser,
		index:     map[tokeniser.Token][]int{},
	}
}

func (r *ReverseIndex) AddDocument(document string, id int) {
	tokens := r.tokeniser.Tokens(document)

	for _, token := range tokens {
		existingIds, ok := r.index[token]

		if !ok {
			r.index[token] = []int{id}
			continue
		}

		r.index[token] = append(existingIds, id)
	}
}

func (r *ReverseIndex) SearchToken(token tokeniser.Token) []int {
	existingsIds, ok := r.index[token]

	if !ok {
		return []int{}
	}

	return existingsIds
}
