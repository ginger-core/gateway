package gateway

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/filter"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

// ProcessFilters processes given request and returns filters
func ProcessFilters(request Request, base query.Query,
	instruction instruction.Instruction) (query.Filter, errors.Error) {
	q, _ := request.GetQuery("filter")
	q2 := q
	if i := len(q2) % 4; i != 0 {
		q2 += strings.Repeat("=", 4-i)
	}
	q3, dErr := base64.StdEncoding.DecodeString(q2)
	if dErr == nil {
		q = string(q3)
	}

	fieldsMap := make(map[string]bool)
	ctx := request.GetContext()
	ctx = context.WithValue(ctx, filter.ProcessedFieldsMapKey, fieldsMap)

	r, err := filter.Parse(ctx, base, instruction, q)
	if err != nil {
		return nil, err
	}
	// defaults
	defaultMatchs := instruction.GetDefaultMatchs()
	for _, m := range defaultMatchs {
		if !fieldsMap[m.Key] {
			r = r.WithMatch(m)
		}
	}
	return r, nil
}
