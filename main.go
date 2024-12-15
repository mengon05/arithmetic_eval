package main

import (
	"fmt"
	"os"

	"github.com/mengon05/arithmetic_eval.git/eval"
	"github.com/mengon05/arithmetic_eval.git/lexer"
)

func main() {
	exp := "(1))+1"
	if len(os.Args) > 1 && os.Args[1] != "" {
		exp = os.Args[1]
	}
	lex := lexer.New()
	tokens := lex.Tokenize(exp)
	evaluator := eval.New(tokens)
	r, err := evaluator.Eval()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s = %d\n", exp, r)
	}
}
