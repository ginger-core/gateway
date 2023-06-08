package instruction

import "github.com/ginger-core/query"

func (i *instruction) WithDefaultMatch(match *query.Match) Instruction {
	if i.defaultMatchs == nil {
		i.defaultMatchs = make([]*query.Match, 0)
	}
	m := i.fieldKeyMap[match.Key]
	if m != nil {
		if h := m.GetCustomHandle(); h != nil {
			match.WithCustomHandle(h)
		}
	}
	i.defaultMatchs = append(i.defaultMatchs, match)
	return i
}

func (i *instruction) GetDefaultMatchs() []*query.Match {
	return i.defaultMatchs
}
