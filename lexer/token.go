package lexer

import (
	"fmt"
)

type Token struct {
	digits int
	Type   TokenType
	Value  int
}

func (lt *Token) AddDigit(digit int) {

	r := lt.Value * 10
	lt.Value = r + digit
	lt.digits++
}
func newTokenWithValue(t TokenType, val int) *Token {
	return &Token{
		Type:   t,
		Value:  val,
		digits: 1,
	}
}
func newToken(t TokenType) *Token {
	return &Token{
		Type: t,
	}
}
func (lt Token) String() string {
	return fmt.Sprintf("{Type %s Value %d}", string(lt.Type), lt.Value)
}
