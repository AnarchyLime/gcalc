package gcalc

import "testing"

func TestInitCalc(t *testing.T) {
	if cs := NewCalc(); cs == nil {
		t.Error("Failed to create a calculator data stack")
	}
}

func TestSingleNumberToContinuousAdd(t *testing.T) {
	cs := NewCalc()
	r, _ := cs.PushKey("4")
	r, _ = cs.PushKey("5")
	t.Log(cs.stack)
	r, _ = cs.PushKey("=")
	t.Log(cs.stack)
	r, _ = cs.PushKey("1")
	r, _ = cs.PushKey("+")
	r, _ = cs.PushKey("2")
	r, _ = cs.PushKey("+")
	r, _ = cs.PushKey("3")
	r, _ = cs.PushKey("=")

	if r != "6" {
		t.Error("Failed to properly reset operations after single entry 6 != " + r)
	}
}

func TestPushKey(t *testing.T) {
	cs := NewCalc()

	cs.PushKey("1")
	r, _ := cs.PushKey("2")
	if r != "12" {
		t.Error("Numbers not representing correctly")
	}

	if r, _ = cs.PushKey("+"); r != "12" {
		t.Error("Operator not returning correct state")
	}

	cs.PushKey("3")
	r, _ = cs.PushKey("4")
	if r != "34" {
		t.Error("Numbers not representing correctly after operator entry")
	}

	if r, _ = cs.PushKey("="); r != "46" {
		t.Error("Result not showing correctly")
	}

	if r, _ = cs.PushKey("5"); r != "5" {
		t.Error("Previous result not clearing properly")
	}

	if r, _ = cs.PushKey("+"); r != "5" {
		t.Error("Previous result not clearing on next operator")
	}

	if r, _ = cs.PushKey("6"); r != "6" {
		t.Error("Previous result not clearing on next operand after next operator")
	}

	if r, _ = cs.PushKey("="); r != "11" {
		t.Error("Previous result not clearing for next operation")
	}
}

func TestPushKeyContinuousOp(t *testing.T) {
	cs := NewCalc()

	cs.PushKey("2")
	cs.PushKey("+")
	cs.PushKey("3")
	cs.PushKey("+")
	cs.PushKey("4")
	cs.PushKey("+")
	cs.PushKey("5")
	r, _ := cs.PushKey("+")

	if r != "14" {
		t.Error("Continuous Add is failing 14 !=", r)
	}
}

func TestPushKeyContinuousEq(t *testing.T) {
	cs := NewCalc()

	cs.PushKey("2")
	cs.PushKey("+")
	cs.PushKey("3")
	cs.PushKey("=")
	cs.PushKey("=")
	cs.PushKey("=")
	r, _ := cs.PushKey("=")

	if r != "14" {
		t.Error("Continuous Equal is failing 14 !=", r)
	}
}
