package tokeniser

import "unicode"

type EnglishTokeniser struct{}

func NewEnglishTokeniser() *EnglishTokeniser {
	return &EnglishTokeniser{}
}

func (t *EnglishTokeniser) Tokens(text string) []Token {
	tokens := []Token{}

	curToken := []rune{}
	for _, ch := range text {
		if unicode.IsSpace(ch) {
			if len(curToken) == 0 {
				continue
			}

			tokens = append(tokens, Token((curToken)))
			curToken = []rune{}
			continue
		}

		curToken = append(curToken, unicode.ToLower(ch))
	}

	if len(curToken) > 0 {
		tokens = append(tokens, Token((curToken)))
	}

	return tokens
}
