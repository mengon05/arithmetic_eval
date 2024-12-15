package eval

import (
	"fmt"

	"github.com/mengon05/arithmetic_eval.git/lexer"
)

type evalTree struct {
	tokens   []*lexer.Token
	position int
	rparam   int
}

func (e *evalTree) Eval() (int, error) {
	n, err := e.exec()
	if err != nil {
		return 0, err
	}
	r, err := n.eval()
	return r, err
}
func New(tokens []*lexer.Token) evalTree {
	return evalTree{
		tokens:   tokens,
		position: -1,
	}
}

func (e *evalTree) next() *lexer.Token {
	e.position++
	if e.position < len(e.tokens) {
		t := e.tokens[e.position]
		return t
	}
	return nil
}
func (e *evalTree) workingToken() *lexer.Token {
	if e.position < len(e.tokens) {
		return e.tokens[e.position]
	}
	return nil
}
func (e *evalTree) exec() (*Node, error) {
	e.position = 0
	return e.level1()

}

// plus and minus
func (e *evalTree) level1() (*Node, error) {
	node, err := e.level2()
	if err != nil {
		return nil, err
	}
	l := e.workingToken()

	for l != nil && (l.Type == '+' || l.Type == '-') {
		e.next()
		tmp := &Node{Token: l}
		tmp.Left = node
		node = tmp
		node.Right, err = e.level2()
		if err != nil {
			return nil, err
		}
		l = e.workingToken()
	}

	return node, nil
}

// mult and div
func (e *evalTree) level2() (*Node, error) {
	node, err := e.level3()
	if err != nil {
		return nil, err
	}
	l := e.workingToken()
	if l != nil && e.rparam == 0 && l.Type == lexer.TokenTypes.RParentesis {
		return nil, fmt.Errorf("unexpected character %c", l.Type)
	}
	for l != nil && (l.Type == '*' || l.Type == '/') {
		e.next()
		tmp := &Node{Token: l}
		tmp.Left = node
		node = tmp
		node.Right, err = e.level3()
		if err != nil {
			return nil, err
		}
		l = e.workingToken()
	}
	return node, nil
}
func (e *evalTree) level3() (*Node, error) {
	wt := e.workingToken()
	if wt == nil {
		return nil, nil
	}
	if wt.Type.IsNumber() {
		e.next()
		return &Node{Token: wt}, nil
	} else if wt.Type == lexer.TokenTypes.LParentesis {
		e.next()
		e.rparam++
		l1, err := e.level1()
		if err != nil {
			return nil, err
		}
		n := e.workingToken()
		if n == nil || n.Type != lexer.TokenTypes.RParentesis {
			return nil, fmt.Errorf("missing right parentesis")
		}
		e.rparam--
		e.next()
		return l1, nil
	} else {
		return nil, fmt.Errorf("unexpected character %c", wt.Type)
	}
}
