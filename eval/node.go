package eval

import "github.com/mengon05/arithmetic_eval.git/lexer"

type Node struct {
	Token *lexer.Token
	Left  *Node
	Right *Node
}

func (n *Node) eval() int {
	a := 0
	b := 0
	if n.Right == nil {
		return n.Token.Value
	} else {
		b = n.Right.eval()
	}
	if n.Left == nil {
		a = n.Token.Value
	} else {
		a = n.Left.eval()
	}
	if fn, ok := opMap[n.Token.Type]; ok {
		return fn(a, b)
	}
	return 0
}

var opMap = map[lexer.TokenType]func(a, b int) int{
	lexer.TokenTypes.Plus: func(a, b int) int {
		return a + b
	},
	lexer.TokenTypes.Minus: func(a, b int) int {
		return a - b
	},
	lexer.TokenTypes.Mult: func(a, b int) int {
		return a * b
	},
	lexer.TokenTypes.Div: func(a, b int) int {
		return a / b
	},
}
