package filter

import (
	"strings"

	"github.com/ginger-core/query"
)

func isOperation(s string) bool {
	switch strings.ToUpper(s) {
	case "AND", "OR":
		return true
	}
	return false
}

func getOperation(op string) query.Operation {
	switch op {
	case "AND":
		return query.OperationAnd
	case "OR":
		return query.OperationOr
	}
	return 0
}
