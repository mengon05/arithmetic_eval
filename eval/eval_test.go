package eval

import (
	"testing"

	"github.com/mengon05/arithmetic_eval.git/lexer"
)

func exec(exp string) int {
	lex := lexer.New()
	tokens := lex.Tokenize(exp)
	evaluator := New(tokens)
	return evaluator.Eval()
}

func assertExpresion(t *testing.T, exp string, val int) {
	result := exec(exp)
	if result != val {
		t.Errorf("%s excpected to %d, but was %d", exp, val, result)
	}
}
func TestEval_goodcases(t *testing.T) {
	assertExpresion(t, "1+1", 2)
	assertExpresion(t, "10+10", 20)
	assertExpresion(t, "100+100", 200)
	assertExpresion(t, "1000+100", 1100)
	assertExpresion(t, "10+1", 11)
	assertExpresion(t, "1*1", 1)
	assertExpresion(t, "2*3", 6)
	assertExpresion(t, "10+(2+8)", 20)
	assertExpresion(t, "10+(2+3+5)", 20)
	assertExpresion(t, "10+(2*5)", 20)
	assertExpresion(t, "1*((1+2)*(2+3))", 15)
	assertExpresion(t, "1*((1+2)*(2+3))/3", 5)
	assertExpresion(t, "2*(3*5)", 30)
	assertExpresion(t, "((1+1)*((1+2)*(2+3)))", 30)
	assertExpresion(t, "((1+1)*((1+2)*(2+3)))/10", 3)
}
