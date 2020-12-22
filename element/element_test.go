package element

import "testing"

func TestGetValue(t *testing.T) {
	element := New(WithKind(OPERATOR), WithValue("+"))
	if element.GetValue() != "+" {
		t.Errorf("GetValue was incorrect, got: %s, want: %s.", element.GetValue(), "+")
	}
}

func TestGetKind(t *testing.T) {
	element := New(WithKind(OPERATOR), WithValue("+"))
	if element.GetKind() != OPERATOR {
		t.Errorf("GetValue was incorrect, got: %d, want: %d.", element.GetKind(), OPERATOR)
	}
}
