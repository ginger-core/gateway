package filter

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
)

// condition is a single boolean condition in a WHERE clause
type condition struct {
	// Key is key of property which will be filled in query
	Key string
	// Operand1 is the left hand side operand
	Operand1 string
	// Operator is e.g. "=", ">"
	Operator query.Operator
	// Operand1 is the right hand side operand
	Operand2 any
	// Operand2IsField determines if Operand2 is a literal or a field name
	Operand2IsField bool
	// CustomHandle reference of field custom handle for custom queries
	CustomHandle query.CustomHandlerFunc
}

func (c *condition) getFilter() (query.Filter, errors.Error) {
	f := query.NewFilter(nil)
	m := &query.Match{
		Key:   c.Key,
		Value: c.Operand2,
	}
	m.Operator = c.Operator
	m = m.WithCustomHandle(c.CustomHandle)
	return f.WithMatch(m), nil
}
