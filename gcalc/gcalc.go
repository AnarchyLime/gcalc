package gcalc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type calcState struct {
	stack       IStringStack
	lastOp      string
	lastOperand string
}

type InputError struct {
	Input string
}

func (e *InputError) Error() string {
	return fmt.Sprintf("Invalid Input: %s", e.Input)
}

func NewCalc() *calcState {
	cs := calcState{new(StringStack), "", "0"}
	cs.stack.Push("0")

	return &cs
}

func (cs *calcState) PushKey(key string) (string, error) {
	switch key {
	case "0":
		fallthrough
	case "1":
		fallthrough
	case "2":
		fallthrough
	case "3":
		fallthrough
	case "4":
		fallthrough
	case "5":
		fallthrough
	case "6":
		fallthrough
	case "7":
		fallthrough
	case "8":
		fallthrough
	case "9":
		return cs.pushDigit(key), nil

	case "+":
		fallthrough
	case "-":
		fallthrough
	case "=":
		return cs.pushOp(key), nil

	default:
		//raise our error!
		var n string
		if n = cs.stack.Top(); !isNumber(n) {
			n = cs.lastOperand
		}
		return n, &InputError{key}
	}
}

func (cs *calcState) ProcessExpr(expr string) (string, error) {
	var result string
	var err error
	r, _ := regexp.Compile(`^[\s\d\+-=]+$`)
	expr = strings.TrimSpace(expr)
	if r.MatchString(expr) {
		fields := strings.Fields(expr)
		for _, f := range fields {
			for _, c := range f {
				result, err = cs.PushKey(string(c))
				if err != nil {
					return "", err
				}
			}
		}
		return result, nil
	} else {
		return "", &InputError{expr}
	}
}

func (cs *calcState) pushDigit(digit string) string {
	top := cs.stack.Top()
	if top == "=" {
		cs.stack.Pop() //pop equal
		cs.stack.Pop() //pop previous result
		cs.stack.Push(digit)
		cs.lastOperand = digit
		return digit
	} else if isOp(top) {
		cs.stack.Push(digit)
		cs.lastOperand = digit
		return digit
	}

	cs.stack.Pop()
	if top == "0" {
		cs.stack.Push(digit)
		cs.lastOperand = digit
	} else {
		cs.stack.Push(top + digit)
		cs.lastOperand = cs.stack.Top()
	}

	return cs.stack.Top()
}

func (cs *calcState) pushOp(op string) string {
	if op == "=" || cs.stack.Count() == 3 {
		return cs.compute(op)
	}

	top := cs.stack.Top()
	if top == "=" {
		cs.stack.Pop()
		top = cs.stack.Top()
	}

	if isOp(top) {
		cs.stack.Pop()
	}

	v := cs.stack.Top()

	cs.stack.Push(op)

	return v
}

func (cs *calcState) compute(op string) string {
	if cs.stack.Count() == 3 {
		b := cs.stack.Pop()
		cop := cs.stack.Pop()
		a := cs.stack.Pop()

		bv, _ := strconv.Atoi(b)
		av, _ := strconv.Atoi(a)

		rv := doOp(av, bv, cop)
		r := strconv.Itoa(rv)

		cs.lastOp = cop

		cs.stack.Push(r)
		if op != ")" {
			cs.stack.Push(op)
		}

		return r
	} else {
		top := cs.stack.Top()
		if top == "=" {
			cs.stack.Pop()
			cs.stack.Push(cs.lastOp)
			cs.stack.Push(cs.lastOperand)
			return cs.compute(op)
		} else if isOp(top) {
			cs.stack.Push(cs.lastOperand)
			return cs.compute(op)
		} else {
			r := cs.stack.Top()
			cs.stack.Push(op)
			return r
		}
	}
}

func isNumber(s string) bool {
	return !isOp(s)
}

func isOp(op string) bool {
	if op == "+" || op == "-" || op == "=" {
		return true
	}
	return false
}

func doOp(a int, b int, op string) int {
	switch op {
	case "-":
		return a - b
	case "+":
		fallthrough
	default:
		return a + b
	}
}
