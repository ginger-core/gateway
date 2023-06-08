package instruction

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
)

type ValidatorFunc func(value any) errors.Error

type Field interface {
	// WithName sets the field name
	WithName(name string) Field
	// GetName returns name of field which
	// client is accessible through this name
	GetName() string
	// WithKey sets the field key, key is the field name
	// which is going to be processed in logic and db
	WithKey(key string) Field
	// GetKey is the key of Field
	// which would be queried in database or processed in logic
	GetKey() string
	// WithSortable enables sortable
	WithSortable() Field
	// IsSortable returns if sort is enabled for current field
	IsSortable() bool
	// WithOperator adds operator to approved opetaions
	WithOperator(operator query.Operator) Field
	// IsOperationEnabled returns availability of given operation
	IsOperationEnabled(operator query.Operator) bool
	// WithValivator adds validator to field
	WithValivator(ValidatorFunc) Field
	// Validate validate given value
	Validate(value any) errors.Error
	// WithGroup set the name of group
	WithGroup(g string) Field
	// Group returns name of group
	Group() string
	// WithCustomHandle to handle specific operation
	WithCustomHandle(f query.CustomHandlerFunc) Field
	GetCustomHandle() query.CustomHandlerFunc
}

type field struct {
	name       string
	key        string
	isSortable bool
	operators  map[query.Operator]bool
	validator  ValidatorFunc
	group      string
	// customHandle to handle specific operation
	customHandle query.CustomHandlerFunc
}

func NewField() Field {
	return new(field)
}

func (f *field) WithName(name string) Field {
	f.name = name
	return f
}

func (f *field) GetName() string {
	return f.name
}

func (f *field) WithKey(key string) Field {
	f.key = key
	return f
}

func (f *field) GetKey() string {
	return f.key
}

func (f *field) WithSortable() Field {
	f.isSortable = true
	return f
}

func (f *field) IsSortable() bool {
	return f.isSortable
}

func (f *field) WithOperator(operator query.Operator) Field {
	if f.operators == nil {
		f.operators = make(map[query.Operator]bool)
	}
	f.operators[operator] = true
	return f
}

func (f *field) IsOperationEnabled(operator query.Operator) bool {
	if f.operators == nil {
		return false
	}
	return f.operators[operator]
}

func (f *field) WithValivator(validator ValidatorFunc) Field {
	f.validator = validator
	return f
}

func (f *field) Validate(value any) errors.Error {
	if f.validator == nil {
		return nil
	}
	return f.validator(value)
}

func (f *field) WithGroup(g string) Field {
	f.group = g
	return f
}

func (f *field) Group() string {
	return f.group
}

func (f *field) WithCustomHandle(_f query.CustomHandlerFunc) Field {
	f.customHandle = _f
	return f
}

func (f *field) GetCustomHandle() query.CustomHandlerFunc {
	return f.customHandle
}
