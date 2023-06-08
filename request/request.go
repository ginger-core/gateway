package request

import (
	"context"

	"github.com/ginger-core/gateway"
	"github.com/google/uuid"
)

type request struct {
	// context as connection context to handle timeout, cancel, etc.
	context context.Context
	// authorization of applicant
	authorization gateway.Authorization
	// id request unique id
	id string
	// query is processed query of request
	query any
	// body is processed body of request
	body any
	// header is processed header of request
	header any
	// hasResponded determines if it has already responded to client or not
	hasResponded bool
	// language current language of request
	language gateway.Language
}

func New() gateway.Request {
	r := new(request)
	uid, _ := uuid.NewRandom()
	r.id = uid.String()
	return r
}
