package main

import (
	"fmt"

	"github.com/mengon05/arithmetic_eval.git/eval"
	"github.com/mengon05/arithmetic_eval.git/lexer"
)

func main() {

	lex := lexer.New()
	tokens := lex.Tokenize("8/2")
	// tokens := lex.Tokenize("69*(4/2)+3*1")
	for _, t := range tokens {
		fmt.Printf("%v\n", t)
	}
	eval.EvalMaster(tokens)
}
