package lexer

type lexer struct {
	patterns []regexPattern
	Tokens []Token
	source string
	pos int
}