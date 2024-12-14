package eval

import (
	"fmt"

	"github.com/mengon05/arithmetic_eval.git/lexer"
)

func EvalMaster(tokens []*lexer.Token) {

	et := New(tokens)
	n := et.Exec()
	r := n.Eval()
	fmt.Printf("%v\n", r)
}

type evalTree struct {
	tokens   []*lexer.Token
	position int
}

func New(tokens []*lexer.Token) evalTree {
	return evalTree{
		tokens:   tokens,
		position: -1,
	}
}

func (e *evalTree) Next() *lexer.Token {
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
func (e *evalTree) Exec() *Node {
	e.position = 0
	return e.level1()

}
func (e *evalTree) level1() *Node {
	node := e.level2()
	l := e.workingToken()
	for l != nil && (l.Type == '+' || l.Type == '-') {
		e.Next()
		tmp := &Node{Token: l}
		tmp.Left = node
		node = tmp
		node.Right = e.level2()
		l = e.workingToken()
	}
	return node
}
func (e *evalTree) level2() *Node {
	node := e.level3()
	l := e.workingToken()
	for l != nil && (l.Type == '*' || l.Type == '/') {
		e.Next()
		tmp := &Node{Token: l}
		tmp.Left = node
		node = tmp
		node.Right = e.level3()
		l = e.workingToken()
	}
	return node
}
func (e *evalTree) level3() *Node {
	wt := e.workingToken()
	if wt == nil {
		return nil
	}
	fmt.Printf("%c = %s\n", wt.Type, string(wt.Type))
	if wt.Type.IsNumber() {
		e.Next()
		return &Node{Token: wt}
	} else {
		fmt.Printf("something wrong %v", wt)
	}
	return nil
}
