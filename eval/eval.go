package eval

import (
	"fmt"

	"github.com/mengon05/arithmetic_eval.git/lexer"
)

func EvalMaster(tokens []*lexer.Token) {

	et := New(tokens)
	r := et.Eval()
	fmt.Printf("%v\n", r)
}

type evalTree struct {
	tokens   []*lexer.Token
	position int
}

func (e *evalTree) Eval() int {
	n := e.exec()
	r := n.eval()
	return r
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
func (e *evalTree) exec() *Node {
	e.position = 0
	return e.level1()

}

// plus and minus
func (e *evalTree) level1() *Node {
	node := e.level2()
	l := e.workingToken()
	for l != nil && (l.Type == '+' || l.Type == '-') {
		e.next()
		tmp := &Node{Token: l}
		tmp.Left = node
		node = tmp
		node.Right = e.level2()
		l = e.workingToken()
	}
	return node
}

// mult and div
func (e *evalTree) level2() *Node {
	node := e.level3()
	l := e.workingToken()
	for l != nil && (l.Type == '*' || l.Type == '/') {
		e.next()
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
	fmt.Printf("%c = %s :: val %d\n", wt.Type, string(wt.Type), wt.Value)
	if wt.Type.IsNumber() {
		e.next()
		return &Node{Token: wt}
	} else if wt.Type == lexer.TokenTypes.LParentesis {
		e.next()
		l1 := e.level1()
		n := e.workingToken()
		if n.Type != lexer.TokenTypes.RParentesis {
			panic(fmt.Sprintf("missing right parentesis, found %c", n.Type))
		}
		e.next()
		return l1
	} else {
		fmt.Printf("something wrong %v", wt)
	}
	return nil
}
