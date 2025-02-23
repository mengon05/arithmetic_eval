package lexer

import (
	"fmt"
	"strconv"
)

type Lexer struct {
}

func New() Lexer {
	return Lexer{}
}

func (l *Lexer) Tokenize(str string) []*Token {
	tokens := []*Token{}
	var lastToken *Token = nil
	for _, r := range str {
		if r >= '0' && r <= '9' {
			digit, err := strconv.Atoi(string(r))
			if err != nil {
				panic(fmt.Sprintf("Not a digit: %c", r))
			}
			if lastToken != nil && lastToken.Type == TokenTypes.Number {
				lastToken.AddDigit(digit)
			} else {
				t := newTokenWithValue(TokenTypes.Number, digit)
				lastToken = t
				tokens = append(tokens, t)
			}
			continue

		}
		if r == '(' {
			if lastToken != nil && lastToken.Type == ')' {
				t := newToken(TokenTypes.Mult)
				lastToken = t
				tokens = append(tokens, t)
			}
		}
		ltt := TokenType(r)
		t := ltt.Token()
		tokens = append(tokens, t)
		lastToken = t

	}
	return tokens
}
