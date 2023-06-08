package sort

import (
	"context"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

var (
	inaccesibleFieldError = errors.Forbidden().
		WithCode(223).
		WithTrace("inaccesible sort field")
)

func Parse(ctx context.Context, instruction instruction.Instruction,
	_query query.Query, queryString string) (query.Sorts, errors.Error) {
	result := query.NewSorts(_query)
	q := strings.TrimSpace(queryString)
	parts := strings.Split(q, ",")
	for _, part := range parts {
		fieldName := part
		if len(fieldName) == 0 {
			continue
		}
		desc := false
		if fieldName[0] == '-' {
			desc = true
			fieldName = fieldName[1:]
		}
		field := instruction.GetField(fieldName)
		if field == nil || !field.IsSortable() {
			return nil, inaccesibleFieldError
		}

		sort := query.NewSort(_query).SortBy(field.GetKey())
		if desc {
			sort.Desc()
		}
		result.WithSort(sort)
	}
	return result, nil
}
