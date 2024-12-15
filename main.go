package main

import (
	"fmt"

	"github.com/mengon05/arithmetic_eval.git/eval"
	"github.com/mengon05/arithmetic_eval.git/lexer"
)

func main() {
	exp := "1*((1+2)*(2+3))"
	lex := lexer.New()
	tokens := lex.Tokenize(exp)
	evaluator := eval.New(tokens)
	r := evaluator.Eval()
	fmt.Printf("%s = %d", exp, r)
}
