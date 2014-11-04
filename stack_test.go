package gcalc

import "testing"

func TestStringStackPush(t *testing.T) {
	a := new(StringStack)
	b := a.Push("Hello")
	a.Push("World")

	if len(*a) != 2 {
		t.Error("did not push values")
	}

	if a != b {
		t.Error("did not return the same reference")
	}
}

func TestStringStackCount(t *testing.T) {
	a := make(StringStack, 2)
	a[0] = "Hello"
	a[1] = "World"

	if a.Count() != 2 {
		t.Error("count failed to return correct value")
	}
}

func TestStringStackPop(t *testing.T) {
	a := new(StringStack)
	a.Push("Hello").Push("World").Push("Foo").Push("Bar")

	s := a.Pop()

	if s != "Bar" {
		t.Error("did not pop correct value")
	}

	if a.Count() != 3 {
		t.Error("did not pop correct number of values")
	}
}

func TestStringStackTop(t *testing.T) {
	a := new(StringStack)
	a.Push("Hello").Push("World").Push("Foo").Push("Bar")

	s := a.Top()

	if s != "Bar" {
		t.Error("Top() did not report top level element")
	}
}
