package instruction

import (
	"github.com/ginger-core/query"
)

type Instruction interface {
	// WithField adds new field to fields
	WithField(f Field) Instruction
	// Map returns all filters
	// map[name]filter
	GetFields() map[string]Field
	// Get filter by name
	GetField(name string) Field
	// WithDefaultSorts set default sort if user did't pass any sort
	WithDefaultSorts(sorts query.Sorts) Instruction
	GetDefaultSorts() query.Sorts
	// WithDefaultMatchs set default match if user did't pass expected one.
	WithDefaultMatch(matchs *query.Match) Instruction
	GetDefaultMatchs() []*query.Match

	WithGroupKeyPrefix(group string, prefix string) Instruction
}

type instruction struct {
	fieldKeyMap   map[string]Field
	fieldNameMap  map[string]Field
	defaultSorts  query.Sorts
	defaultMatchs []*query.Match
}

func NewInstruction() Instruction {
	return new(instruction)
}

func (i *instruction) WithGroupKeyPrefix(
	group string, prefix string) Instruction {
	for _, f := range i.fieldNameMap {
		if f.Group() == group {
			f.WithKey(prefix + f.GetKey())
		}
	}
	return i
}
