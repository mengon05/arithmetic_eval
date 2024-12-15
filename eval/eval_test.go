package eval

import (
	"testing"

	"github.com/mengon05/arithmetic_eval.git/lexer"
)

func assertError(t *testing.T, exp string, msg string) {
	_, err := exec(exp)
	if err == nil {
		t.Errorf("exp %s does not return error", exp)
		return
	}
	if err.Error() != msg {
		t.Errorf("Error does not match expected \"%s\", actual \"%s\"", msg, err.Error())
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
func TestEval_invalid_expresions(t *testing.T) {
	assertError(t, "a+b", "unexpected character a")
	assertError(t, "1+b", "unexpected character b")
	assertError(t, "1/0", "invalid operation: 1 divided by 0")
	assertError(t, "100/(3-3)", "invalid operation: 100 divided by 0")
}

func TestEval_parentesis_error(t *testing.T) {
	assertError(t, "(1+1", "missing right parentesis")
	assertError(t, "(1+(1)", "missing right parentesis")
	assertError(t, ")", "unexpected character )")
	assertError(t, "1)", "unexpected character )")
	assertError(t, "(1+1))", "unexpected character )")
	assertError(t, "(1))+1", "unexpected character )")
	assertError(t, "()", "unexpected character )")
	assertError(t, "(()", "unexpected character )")
	assertError(t, "1+(())", "unexpected character )")
}

func exec(exp string) (int, error) {
	lex := lexer.New()
	tokens := lex.Tokenize(exp)
	evaluator := New(tokens)
	return evaluator.Eval()
}

func assertExpresion(t *testing.T, exp string, val int) {
	result, err := exec(exp)
	if err != nil {
		t.Errorf("Unexpected error evaluating %s: %s", exp, err.Error())
	}
	if result != val {
		t.Errorf("%s expected \"%d\", actual \"%d\"", exp, val, result)
	}
}
