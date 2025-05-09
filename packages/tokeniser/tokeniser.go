package tokeniser

type Token string

type Tokeniser interface {
	Tokens(text string) []Token
}
