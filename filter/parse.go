package filter

import (
	"context"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

var (
	inaccesibleFilterError = errors.Forbidden().
				WithCode(222).
				WithTrace("inaccesible filter")
	inaccesibleOperationError = errors.Forbidden().
					WithCode(226).
					WithTrace("inaccesible operation")
)

func (p *parser) parse(ctx context.Context, base query.Query,
	instruction instruction.Instruction) (query.Filter, errors.Error) {
	q, err := p.doParse(ctx, base, instruction)
	p.err = err
	if p.err == nil {
		p.err = p.validate()
	}
	if p.err != nil {
		p.err = errors.Validation(err).
			WithTrace("p.doParse")
	}
	return q, p.err
}

func (p *parser) doParse(ctx context.Context, base query.Query,
	instruction instruction.Instruction) (query.Filter, errors.Error) {
	fields := ctx.Value(ProcessedFieldsMapKey).(map[string]bool)
	lastOperation := ""
	var currentFilter query.Filter = query.NewFilter(base)
	for {
		if p.i >= len(p.queryString) {
			// done
			return currentFilter, nil
		}
		part := p.peek()

		switch part {
		case "(":
			// group start
			p.goupStarted += 1
			p.step = stepWhereField
			p.pop()
			f, err := p.doParse(ctx, base, instruction)
			if err != nil {
				return nil, err
			}
			if err = p.pushFilter(currentFilter, lastOperation, f); err != nil {
				return nil, err
			}
			continue
		case ")":
			// done
			// group end
			if p.goupStarted == 0 {
				return nil, errors.Validation().
					WithTrace("handleFilter.!groupStarted")
			}
			p.goupStarted -= 0
			p.pop()
			return currentFilter, nil
		}

		switch p.step {
		case stepWhereField:
			cond, err := p.getCondition(instruction)
			if err != nil {
				return nil, err
			}
			fields[cond.Key] = true
			filter2, err := cond.getFilter()
			if err != nil {
				return nil, err
			}
			if currentFilter == nil {
				// is new
				currentFilter = filter2
				continue
			}
			err = p.pushFilter(currentFilter, lastOperation, filter2)
			if err != nil {
				return nil, err
			}
		case stepWhereMatchDone:
			// match done
			if isOperation(part) {
				p.step = stepOperation
				continue
			}
			p.pop()
		case stepOperation:
			lastOperation = strings.ToUpper(part)
			p.step = stepWhereField
			p.pop()
		}
	}
}

func (p *parser) getCondition(
	instruction instruction.Instruction) (*condition, errors.Error) {
	var cond *condition
	for {
		if p.i >= len(p.queryString) {
			break
		}
		part := p.peek()
		switch p.step {
		case stepWhereField:
			if !isIdentifier(part) {
				return nil, errors.Validation().
					WithTrace("doParse.!isIdentifier").
					WithDesc("at WHERE: expected field")
			}
			field := instruction.GetField(part)
			if field == nil {
				return nil, inaccesibleFilterError
			}
			cond = &condition{
				Key:          field.GetKey(),
				Operand1:     part,
				CustomHandle: field.GetCustomHandle(),
			}
			p.pop()
			p.step = stepWhereOperator
		case stepWhereOperator:
			switch part {
			case "=":
				cond.Operator = query.Equal
			case "!=":
				cond.Operator = query.NotEqual
			case ">":
				cond.Operator = query.Greater
			case ">=":
				cond.Operator = query.GreaterEqual
			case "<":
				cond.Operator = query.Lower
			case "<=":
				cond.Operator = query.LowerEqual
			case "&=":
				cond.Operator = query.BitwiseAndEqual
			case "&!=":
				cond.Operator = query.BitwiseAndNotEqual
			case "&~":
				cond.Operator = query.BitwiseIs
			case "&!~":
				cond.Operator = query.BitwiseIsNot
			case "IS":
				cond.Operator = query.Is
			case "IN":
				cond.Operator = query.In
			default:
				return nil, errors.Validation().
					WithTrace("doParse.operator.unknown").
					WithDesc("unknown operator")
			}
			// validate
			field := instruction.GetField(cond.Operand1)
			if !field.IsOperationEnabled(cond.Operator) {
				return nil, inaccesibleOperationError
			}
			//
			p.pop()
			p.step = stepWhereValue
		case stepWhereValue:
			if isIdentifier(part) {
				cond.Operand2 = part
				cond.Operand2IsField = true
			} else {
				quotedValue, ln := p.peekQuotedStringWithLength()
				if ln == 0 {
					return nil, errors.Validation().
						WithTrace("doParse.peekQuotedStringWithLength.ln=0").
						WithDesc("expected quoted value")
				}
				cond.Operand2 = quotedValue
				cond.Operand2IsField = false

				switch cond.Operator {
				case query.In:
					parts := strings.Split(strings.Trim(quotedValue, "()"), ",")
					cond.Operand2 = parts
				}
			}
			if cond.Operand2 == "null" || cond.Operand2 == "NULL" {
				cond.Operand2 = nil
			}
			// validate
			field := instruction.GetField(cond.Operand1)
			if err := field.Validate(cond.Operand2); err != nil {
				return nil, err.
					WithTrace("doParse.stepWhereValue.Validate")
			}
			// normalize
			switch cond.Operator {
			case query.In:
				if s, ok := cond.Operand2.(string); ok {
					cond.Operand2 = strings.Split(s, ",")
				}
			}
			//
			p.pop()
			p.step = stepWhereMatchDone
			// match done
			return cond, nil
		}
	}
	return nil, errors.New().
		WithTrace("getCondition.notFound")
}

func (p *parser) pushFilter(
	filter query.Filter, op string, filter2 query.Filter) errors.Error {
	if len(filter.GetOperations()) == 0 {
		filter.WithAnd(filter2)
		return nil
	}
	switch op {
	case "AND":
		filter.WithAnd(filter2)
	case "OR":
		filter.WithOr(filter2)
	default:
		filter.WithBase(filter2)
	}
	return nil
}
