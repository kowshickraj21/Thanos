package main

import (
	"regexp"
	"strings"
)

type TokenType int

type Token struct {
	Type  string
	Value string
}

var keywords = map[string]struct{}{
	"if": {}, "else": {}, "for": {}, "while": {}, "return": {},
}

var operators = map[string]struct{}{
	"+": {}, "-": {}, "*": {}, "/": {}, "=": {},
}

func Tokenize(program string) []Token {
	var tokens []Token
	var word strings.Builder
	i := 0
	for i < len(program) {
		ch := program[i]

		if regexp.MustCompile(`^[a-zA-Z]$`).MatchString(string(ch)) {
			for i < len(program) && regexp.MustCompile(`^[a-zA-Z0-9]$`).MatchString(string(program[i])) {
				word.WriteByte(program[i])
				i++
			}

			lexeme := word.String()
			word.Reset()

			if _, exists := keywords[lexeme]; exists {
				tokens = append(tokens, Token{Type: "Keyword", Value: lexeme})
			} else {
				tokens = append(tokens, Token{Type: "Identifier", Value: lexeme})
			}
			continue
		}

		if regexp.MustCompile(`^[0-9]$`).MatchString(string(ch)) {
			for i < len(program) && regexp.MustCompile(`^[0-9]$`).MatchString(string(program[i])) {
				word.WriteByte(program[i])
				i++
			}
			tokens = append(tokens, Token{Type: "IntConstant", Value: word.String()})
			word.Reset()
			continue
		}

		if ch == ';' {
			tokens = append(tokens, Token{Type: "Semicolon", Value: string(ch)})
		} else if _, exists := operators[string(ch)]; exists {
			tokens = append(tokens, Token{Type: "Operator", Value: string(ch)})
		}
		i++
	}

	return tokens
}
