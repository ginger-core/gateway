package gateway

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

type Request interface {
	WithContext(ctx context.Context) Request
	GetContext() context.Context

	WithConn(conn any) Request
	GetConn() any

	GetHeader(key string) string
	GetParam(key string) string
	GetQuery(key string) (string, bool)

	IsAuthenticated() bool
	SetAuthorization(authorization Authorization)
	WithAuthorization(authorization Authorization) Request
	GetAuthorization() Authorization

	WithId(id string) Request
	GetId() string

	WithLanguage(language Language) Request
	GetLanguage() Language

	ProcessBody(ref any) errors.Error
	GetBody() any
	// ProcessFilters processes filters and update given base query and
	// finally returns new one in result
	ProcessFilters(base query.Query,
		instruction instruction.Instruction) (query.Query, errors.Error)
	// ProcessSort processes sort in request and
	// updates query and returns new one in result
	ProcessSort(q query.Query,
		instruction instruction.Instruction) (query.Query, errors.Error)
	// GetPagination returns size and page number
	// returns defaults if not passed
	ProcessPagination(q query.Query,
		instruction instruction.Instruction) query.Query

	ProcessQueries(ref any) errors.Error
	GetQueries() any

	ProcessHeaders(ref any) errors.Error
	GetHeaders() any

	SetResponded()
	HasResponded() bool
}
