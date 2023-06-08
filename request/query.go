package request

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

func (r *request) ProcessFilters(base query.Query,
	instruction instruction.Instruction) (query.Query, errors.Error) {
	filters, err := gateway.ProcessFilters(r, base, instruction)
	if err != nil {
		return nil, err
	}
	if filters == nil {
		return base, nil
	}

	return filters, nil
}

func (r *request) GetQuery(key string) (string, bool) {
	return "", false
}

func (r *request) ProcessQueries(ref any) errors.Error {
	return nil
}

func (r *request) GetQueries() any {
	return r.query
}
