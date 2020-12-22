package element

// Option is a closure that is used for accessing the local variables.
type Option func(element *Element)

// WithValue is used for initialize field value.
func WithValue(value string) Option {
	return func(element *Element) {
		element.value = value
	}
}

// WithKind is used for initialize field kind.
func WithKind(kind int) Option {
	return func(element *Element) {
		element.kind = kind
	}
}
