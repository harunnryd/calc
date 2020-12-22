package scanner

import (
	"bytes"
	"github.com/harunnryd/calc/element"
	"github.com/harunnryd/calc/stack"
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const (
	NEGATIVE = "-1"
)

// IScanner is an interface that stores the methods that Scanner struct will use.
type IScanner interface {
	// DoScan is used for splitting string to the element.Element of stack.Stack.
	// It returns any errors written.
	DoScan() (err error)

	// IsOperator is used for checking rune item.
	// It returns bool.
	IsOperator(r rune) bool

	// DoSolve is used for solving the problem and returning the result.
	// It returns string.
	DoSolve() string
}

// Scanner is an struct that implements IScanner methods.
type Scanner struct {
	quote string
	stack stack.IStack
}

// New it returns instance of Scanner that implements IScanner methods.
func New(opts ...Option) IScanner {
	scanner := new(Scanner)
	for _, opt := range opts {
		opt(scanner)
	}
	return scanner
}

// DoScan is used for splitting string to the element.Element of stack.Stack.
// It returns any errors written.
func (scanner *Scanner) DoScan() (err error) {
	splitter := func(r rune) bool {
		// Tricks: when the first character is operand negative,
		// Push operand negative and operator multiply to the Stack.
		if scanner.IsOperator(r) && scanner.stack.Peek().GetKind() == element.ERROR {
			scanner.stack.Push(element.New(
				element.WithValue(NEGATIVE),
				element.WithKind(element.OPERAND),
			))

			scanner.stack.Push(element.New(
				element.WithValue("*"),
				element.WithKind(element.OPERATOR),
			))
		} else if scanner.IsOperator(r) && scanner.stack.Peek().GetKind() == element.OPERATOR {
			currentStack := scanner.stack.Pop()
			if currentStack.GetValue() == string(r) {
				scanner.stack.Push(element.New(
					element.WithValue("+"),
					element.WithKind(element.OPERATOR),
				))
			}
			if currentStack.GetValue() != string(r) {
				scanner.stack.Push(element.New(
					element.WithValue("-"),
					element.WithKind(element.OPERATOR),
				))
			}
		} else if unicode.IsDigit(r) && scanner.stack.Peek().GetKind() == element.OPERAND {
			var b bytes.Buffer
			b.WriteString(scanner.stack.Pop().GetValue())
			b.WriteString(string(r))

			scanner.stack.Push(element.New(
				element.WithValue(b.String()),
				element.WithKind(element.OPERAND),
			))
		} else if unicode.IsDigit(r) {
			scanner.stack.Push(element.New(
				element.WithValue(string(r)),
				element.WithKind(element.OPERAND),
			))
		} else if scanner.IsOperator(r) {
			scanner.stack.Push(element.New(
				element.WithValue(string(r)),
				element.WithKind(element.OPERATOR),
			))
		} else if unicode.IsLetter(r) {
			err = errors.New("validation error")
		}
		return true
	}

	strings.FieldsFunc(scanner.quote, splitter)

	return
}

// IsOperator is used for checking rune item.
// It returns bool.
func (scanner *Scanner) IsOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

var operatorData = map[string]struct {
	fx func(x, y float64) float64
}{
	"+": {func(x, y float64) float64 { return x + y }},
	"-": {func(x, y float64) float64 { return x - y }},
	"*": {func(x, y float64) float64 { return x * y }},
}

// DoSolve is used for solving the problem and returning the result.
// It returns string.
func (scanner *Scanner) DoSolve() string {
	stackOperand := stack.New()
	stackOperator := stack.New()
	currentStack := scanner.stack

	for _, ele := range currentStack.GetElements() {
		switch ele.GetKind() {
		case 1:
			stackOperand.Push(ele)
		case 2:
			stackOperator.Push(ele)
		}
	}

	for _, v := range stackOperator.GetElements() {
		switch v.GetKind() {
		case 2:
			x, _ := strconv.ParseFloat(stackOperand.Shift().GetValue(), 64)
			y, _ := strconv.ParseFloat(stackOperand.Shift().GetValue(), 64)

			f := operatorData[stackOperator.Shift().GetValue()].fx
			result := f(x, y)

			stackOperand.Prepend(element.New(
				element.WithValue(strconv.FormatFloat(result, 'f', -1, 64)),
				element.WithKind(element.OPERAND),
			))
		}
	}

	return stackOperand.Pull(0).GetValue()
}
