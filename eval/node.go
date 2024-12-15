package eval

import (
	"fmt"

	"github.com/mengon05/arithmetic_eval.git/lexer"
)

type Node struct {
	Token *lexer.Token
	Left  *Node
	Right *Node
}

func (n *Node) eval() (int, error) {
	a := 0
	b := 0
	var err error
	if n.Right == nil {
		return n.Token.Value, nil
	} else {
		b, err = n.Right.eval()
		if err != nil {
			return 0, err
		}
	}
	if n.Left == nil {
		a = n.Token.Value
	} else {
		a, err = n.Left.eval()
		if err != nil {
			return 0, err
		}
	}
	if fn, ok := opMap[n.Token.Type]; ok {
		return fn(a, b)
	}
	return 0, fmt.Errorf("unsuppoted operator %c", n.Token.Type)
}

var opMap = map[lexer.TokenType]func(a, b int) (int, error){
	lexer.TokenTypes.Plus: func(a, b int) (int, error) {
		return a + b, nil
	},
	lexer.TokenTypes.Minus: func(a, b int) (int, error) {
		return a - b, nil
	},
	lexer.TokenTypes.Mult: func(a, b int) (int, error) {
		return a * b, nil
	},
	lexer.TokenTypes.Div: func(a, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("invalid operation: %d divided by %d", a, b)
		}
		return a / b, nil
	},
}
