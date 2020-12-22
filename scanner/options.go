package scanner

import "github.com/harunnryd/calc/stack"

// Option is a closure that is used for accessing the local variables.
type Option func(scanner *Scanner)

// WithQuote ...
func WithQuote(quote string) Option {
	return func(scanner *Scanner) {
		scanner.quote = quote
	}
}

// WithStack ...
func WithStack(stack stack.IStack) Option {
	return func(scanner *Scanner) {
		scanner.stack = stack
	}
}
