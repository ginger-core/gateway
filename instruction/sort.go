package instruction

import "github.com/ginger-core/query"

func (i *instruction) WithDefaultSorts(sorts query.Sorts) Instruction {
	i.defaultSorts = sorts
	return i
}

func (i *instruction) GetDefaultSorts() query.Sorts {
	if i.defaultSorts == nil {
		return nil
	}
	return i.defaultSorts.Clone()
}
