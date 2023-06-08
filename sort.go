package gateway

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/gateway/sort"
	"github.com/ginger-core/query"
)

// ProcessSort processes given request and returns filters
func ProcessSort(request Request, query query.Query,
	instruction instruction.Instruction) (query.Sorts, errors.Error) {
	q, ok := request.GetQuery("sort")
	if !ok {
		return nil, nil
	}
	if q == "" {
		return nil, nil
	}
	r, err := sort.Parse(request.GetContext(), instruction, query, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}
