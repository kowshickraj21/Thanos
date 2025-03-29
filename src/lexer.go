package main

import (
	"errors"
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

var datatypes = map[string]struct{}{
	"int": {}, "char": {}, "long": {}, "float": {}, "bool": {},
}

var comparators = map[string]struct{}{
	"<": {}, ">": {}, "==": {}, "!=": {}, "<=":{}, ">=": {},
}

var assignments = map[string]struct{}{
	"=": {}, "+=": {}, "-=": {}, "*=": {}, "/=": {}, "%=": {},
}

var arithmetics = map[string]struct{}{
	"+": {}, "-": {}, "*": {}, "/": {}, "%": {},
}

var symbols = map[string]struct{}{
	";": {}, "(": {}, ")": {}, "{": {}, "}":{},
}

func Tokenize(program string) ([]Token,error) {
	var tokens []Token
	var word strings.Builder
	i := 0
	for i < len(program) {
		ch := program[i]
		if(string(ch) == " " || ch == 13 || ch == 10){
			i++;
			continue;
		}
		if regexp.MustCompile(`^[a-zA-Z]$`).MatchString(string(ch)) {
			for i < len(program) && regexp.MustCompile(`^[a-zA-Z0-9]$`).MatchString(string(program[i])) {
				word.WriteByte(program[i])
				i++
			}

			lexeme := word.String()
			word.Reset()

			if _, exists := keywords[lexeme]; exists {
				tokens = append(tokens, Token{Type: "Keyword", Value: lexeme})
			} else if _, exists := datatypes[lexeme]; exists {
				tokens = append(tokens, Token{Type: "DataType", Value: lexeme})
			} else {
				tokens = append(tokens, Token{Type: "Identifier", Value: lexeme})
			}
			continue
		}

		if regexp.MustCompile(`^[0-9]$`).MatchString(string(ch)) {
			for i < len(program) && regexp.MustCompile(`^[0-9]$`).MatchString(string(program[i])) {
				word.WriteByte(program[i])
				i++;
			}
			tokens = append(tokens, Token{Type: "IntConstant", Value: word.String()})
			word.Reset()
			continue
		}

		if i+1 < len(program) {
			operator := string(program[i]) + string(program[i+1]);
			if _, exists := comparators[operator]; exists {
				tokens = append(tokens, Token{Type: "Comparator", Value: operator});
				i += 2;
				continue;
			} else if _, exists := assignments[operator]; exists {
				tokens = append(tokens, Token{Type: "Assignment", Value: operator})
				i += 2;
				continue; 
			}
		}

		if _, exists := comparators[string(ch)]; exists {
			tokens = append(tokens, Token{Type: "Comparator", Value: string(ch)});
			i ++;
			continue;
		} else if _, exists := assignments[string(ch)]; exists {
			tokens = append(tokens, Token{Type: "Assignments", Value: string(ch)})
			i ++;
			continue; 
		} else if _, exists := arithmetics[string(ch)]; exists {
			tokens = append(tokens, Token{Type: "Arithmetics", Value: string(ch)});
			i ++;
			continue;
		} else if _, exists := symbols[string(ch)]; exists {
			tokens = append(tokens, Token{Type: "Symbol", Value: string(ch)})
			i ++;
			continue; 
		}
		return tokens, errors.New("Not a valid Syntax");
	}

	return tokens,nil
}
