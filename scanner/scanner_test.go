package scanner

import (
	"github.com/harunnryd/calc/stack"
	"testing"
)

func TestIsOperator(t *testing.T) {
	nScanner := New()
	if !nScanner.IsOperator('+') {
		t.Errorf("IsOperator was incorrect, got: %t, want: %t.", nScanner.IsOperator('+'), !nScanner.IsOperator('+'))
	}

	if nScanner.IsOperator('2') {
		t.Errorf("IsOperator was incorrect, got: %t, want: %t.", !nScanner.IsOperator('2'), nScanner.IsOperator('2'))
	}
}

func TestDoScan(t *testing.T) {
	nScanner := New(
		WithQuote("-12 + -19 - -2 + 10 + -50"),
		WithStack(stack.New()),
	)

	if err := nScanner.DoScan(); err != nil {
		t.Errorf("DoScan was incorrect, got: %v, want: %v.", err, nil)
	}

	nScanner = New(
		WithQuote("asda12 + -19 - -2 + 10 + -50"),
		WithStack(stack.New()),
	)

	if err := nScanner.DoScan(); err != nil {
		if err.Error() != "validation error" {
			t.Errorf("DoScan was incorrect, got: %v, want: %v.", err.Error(), "validation error")
		}
	}
}

func TestDoSolve(t *testing.T) {
	nScanner := New(
		WithQuote("-12 + -19 - -2 + 10 + -50"),
		WithStack(stack.New()),
	)

	if err := nScanner.DoScan(); err != nil {
		t.Errorf("DoSolve was incorrect, got: %v, want: %v.", err, nil)
	}

	if result := nScanner.DoSolve(); result != "-69" {
		t.Errorf("DoSolve was incorrect, got: %s, want: %s.", result, "-69")
	}
}
