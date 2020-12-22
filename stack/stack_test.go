package stack

import (
	"calc/element"
	"testing"
)

func TestPush(t *testing.T) {
	nStack := New()
	nStack.Push(
		element.New(element.WithValue("-"), element.WithKind(element.OPERATOR)),
		element.New(element.WithValue("12"), element.WithKind(element.OPERAND)),
	)

	if len(nStack.GetElements()) != 2 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 2)
	}

	if nStack.Peek().GetKind() != element.OPERAND {
		t.Errorf("Push was incorrect, got: %d, want: %d.", nStack.Peek().GetKind(), element.OPERAND)
	}
}

func TestPrepend(t *testing.T) {
	nStack := New()
	nStack.Prepend(element.New(element.WithValue("-"), element.WithKind(element.OPERATOR)))
	nStack.Prepend(element.New(element.WithValue("12"), element.WithKind(element.OPERAND)))

	if len(nStack.GetElements()) != 2 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 2)
	}

	if currentKind := nStack.Shift().GetKind(); currentKind != element.OPERAND {
		t.Errorf("Push was incorrect, got: %d, want: %d.", currentKind, element.OPERAND)
	}
}

func TestPop(t *testing.T) {
	nStack := New()
	nStack.Push(
		element.New(element.WithValue("-"), element.WithKind(element.OPERATOR)),
		element.New(element.WithValue("12"), element.WithKind(element.OPERAND)),
	)

	if len(nStack.GetElements()) != 2 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 2)
	}

	if currentKind := nStack.Pop().GetKind(); currentKind != element.OPERAND {
		t.Errorf("Push was incorrect, got: %d, want: %d.", currentKind, element.OPERAND)
	}
}

func TestShift(t *testing.T) {
	nStack := New()
	nStack.Push(element.New(element.WithValue("12"), element.WithKind(element.OPERAND)))
	nStack.Prepend(element.New(element.WithValue("-"), element.WithKind(element.OPERATOR)))
	nStack.Prepend(element.New(element.WithValue("12"), element.WithKind(element.OPERAND)))

	if len(nStack.GetElements()) != 3 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 3)
	}

	currentStack := nStack.Shift()

	if currentStack.GetKind() != element.OPERAND {
		t.Errorf("Push was incorrect, got: %d, want: %d.", currentStack.GetKind(), element.OPERAND)
	}

	if currentStack.GetValue() != "12" {
		t.Errorf("Push was incorrect, got: %s, want: %s.", currentStack.GetValue(), "12")
	}
}

func TestPeek(t *testing.T) {
	nStack := New()
	nStack.Push(element.New(element.WithValue("10"), element.WithKind(element.OPERAND)))
	nStack.Prepend(element.New(element.WithValue("-"), element.WithKind(element.OPERATOR)))
	nStack.Prepend(element.New(element.WithValue("15"), element.WithKind(element.OPERAND)))

	if len(nStack.GetElements()) != 3 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 3)
	}

	currentStack := nStack.Peek()

	if currentStack.GetKind() != element.OPERAND {
		t.Errorf("Push was incorrect, got: %d, want: %d.", currentStack.GetKind(), element.OPERAND)
	}

	if currentStack.GetValue() != "10" {
		t.Errorf("Push was incorrect, got: %s, want: %s.", currentStack.GetValue(), "10")
	}
}

func TestPull(t *testing.T) {
	nStack := New()
	nStack.Push(element.New(element.WithValue("10"), element.WithKind(element.OPERAND)))
	nStack.Prepend(element.New(element.WithValue("-"), element.WithKind(element.OPERATOR)))
	nStack.Prepend(element.New(element.WithValue("15"), element.WithKind(element.OPERAND)))

	if len(nStack.GetElements()) != 3 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 3)
	}

	currentStack := nStack.Pull(1)

	if currentStack.GetKind() != element.OPERATOR {
		t.Errorf("Push was incorrect, got: %d, want: %d.", currentStack.GetKind(), element.OPERATOR)
	}

	if currentStack.GetValue() != "-" {
		t.Errorf("Push was incorrect, got: %s, want: %s.", currentStack.GetValue(), "-")
	}
}

func TestGetElements(t *testing.T) {
	nStack := New()
	nStack.Push(element.New(element.WithValue("10"), element.WithKind(element.OPERAND)))
	nStack.Prepend(element.New(element.WithValue("-"), element.WithKind(element.OPERATOR)))
	nStack.Prepend(element.New(element.WithValue("15"), element.WithKind(element.OPERAND)))

	if len(nStack.GetElements()) != 3 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 3)
	}

	nStack.Pop()
	nStack.Prepend(element.New(element.WithValue("-"), element.WithKind(element.OPERATOR)))

	if len(nStack.GetElements()) != 3 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 3)
	}

	nStack.Shift()

	if len(nStack.GetElements()) != 2 {
		t.Errorf("Push was incorrect, got: %d, want: %d.", len(nStack.GetElements()), 2)
	}
}
