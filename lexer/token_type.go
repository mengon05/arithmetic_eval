package lexer

type TokenType rune

var TokenTypes = struct {
	Number      TokenType
	LParentesis TokenType
	RParentesis TokenType
	Plus        TokenType
	Minus       TokenType
	Mult        TokenType
	Div         TokenType
}{
	Number:      '0',
	LParentesis: '(',
	RParentesis: ')',
	Plus:        '+',
	Minus:       '-',
	Mult:        '*',
	Div:         '/',
}

func (ltt TokenType) Token() *Token {
	return newToken(ltt)
}

func (ltt TokenType) IsOperator() bool {
	if (ltt >= '0' && ltt <= '9') || ltt == '(' || ltt == ')' {
		return false
	} else {
		return true
	}
}
func (ltt TokenType) IsNumber() bool {
	return ltt >= '0' && ltt <= '9'
}
