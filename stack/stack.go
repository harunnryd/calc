package stack

import "github.com/harunnryd/calc/element"

// IStack is an interface that stores the methods that Stack struct will use.
type IStack interface {
	// Push is used for appends an item to the end of the element.
	Push(e ...element.IElement)

	// Prepend is used for appends an item to the beginning of the element.
	Prepend(e ...element.IElement)

	// Pop is used for removes and getting the last item from the element.
	// It returns element.
	Pop() (e element.IElement)

	// Shift is used for removes and getting the first item from the element.
	// It returns element.
	Shift() (e element.IElement)

	// Peek is used for getting last item from the element.
	// It returns element.
	Peek() (e element.IElement)

	// Pull is used for removes and getting the first item from the element.
	//It returns element.
	Pull(n int) (e element.IElement)

	// GetElements is used for getting array of element.IElement.
	// It returns element.
	GetElements() (e []element.IElement)
}

// Stack is an struct that implements IStack methods.
type Stack struct {
	Elements []element.IElement
}

// New it returns instance of Stack that implements IStack methods.
func New(opts ...Option) IStack {
	stack := new(Stack)
	for _, opt := range opts {
		opt(stack)
	}
	return stack
}

// Push is used for appends an item to the end of the element.
func (s *Stack) Push(e ...element.IElement) {
	s.Elements = append(s.Elements, e...)
}

// Prepend is used for appends an item to the beginning of the element.
func (s *Stack) Prepend(e ...element.IElement) {
	s.Elements = append(e, s.Elements...)
}

// Pop is used for removes and getting the last item from the element.
// It returns element.
func (s *Stack) Pop() (e element.IElement) {
	if len(s.Elements) == 0 {
		e = element.New()
		return
	}

	e = s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return
}

// Shift is used for removes and getting the first item from the element.
// It returns element.
func (s *Stack) Shift() (e element.IElement) {
	if len(s.Elements) == 0 {
		e = element.New()
		return
	}

	e = s.Elements[0]
	s.Elements = s.Elements[1:]
	return
}

// Peek is used for getting last item from the element.
// It returns element.
func (s *Stack) Peek() (e element.IElement) {
	if len(s.Elements) == 0 {
		e = element.New()
		return
	}

	return s.Elements[len(s.Elements)-1]
}

// Pull is used for removes and getting the first item from the element.
//It returns element.
func (s *Stack) Pull(n int) (e element.IElement) {
	if len(s.Elements) == 0 || len(s.Elements) < n {
		e = element.New()
		return
	}

	return s.Elements[n]
}

// GetElements is used for getting array of element.IElement.
// It returns element.
func (s *Stack) GetElements() (e []element.IElement) {
	return s.Elements
}
