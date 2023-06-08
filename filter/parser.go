package filter

import (
	"context"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

func Parse(ctx context.Context, base query.Query,
	instruction instruction.Instruction,
	queryString string) (query.Filter, errors.Error) {
	p := &parser{
		0,
		strings.TrimSpace(queryString),
		stepWhereField,
		nil,
		0,
	}
	q, err := p.parse(ctx, base, instruction)
	if err != nil {
		return nil, err
	}
	return q, nil
}

type parser struct {
	i           int
	queryString string
	step        step
	err         errors.Error
	goupStarted int
}
