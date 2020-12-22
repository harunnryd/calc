package element

const (
	ERROR = iota
	OPERAND
	OPERATOR
)

// IElement is an interface that stores the methods that Element struct will use.
type IElement interface {
	// GetValue is used for getting field value.
	// It returns string.
	GetValue() string

	// GetKind is used for getting field kind.
	// It returns int.
	GetKind() int
}

// Element is an struct that implements IElement methods.
type Element struct {
	value string
	kind  int
}

// New it returns instance of Element that implements IElement methods.
func New(opts ...Option) IElement {
	element := new(Element)
	for _, opt := range opts {
		opt(element)
	}
	return element
}

// GetValue is used for getting field value.
// It returns string.
func (element *Element) GetValue() string {
	return element.value
}

// GetKind is used for getting field kind.
// It returns int.
func (element *Element) GetKind() int {
	return element.kind
}
