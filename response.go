package gateway

import "github.com/ginger-core/errors"

type Responder interface {
	Respond(request Request, status Status, response any)
	RespondError(request Request, err errors.Error)
}

type Response interface {
	WithStatus(status Status) Response
	GetStatus() Status

	WithContentType(t ContentType) Response
	GetContentType() ContentType

	WithHeader(key, value string) Response
	GetHeaders() map[string]string

	WithBody(body Body) Response
	GetBody() Body
}

type response struct {
	status      Status
	contentType ContentType
	headers     map[string]string
	body        Body
}

func NewResponse() Response {
	return &response{
		headers: make(map[string]string),
	}
}

func (r *response) WithStatus(status Status) Response {
	r.status = status
	return r
}

func (r *response) GetStatus() Status {
	return r.status
}

func (r *response) WithContentType(t ContentType) Response {
	r.contentType = t
	return r
}

func (r *response) GetContentType() ContentType {
	return r.contentType
}

func (r *response) WithHeader(key, value string) Response {
	r.headers[key] = value
	return r
}

func (r *response) GetHeaders() map[string]string {
	return r.headers
}

func (r *response) WithBody(body Body) Response {
	r.body = body
	return r
}

func (r *response) GetBody() Body {
	return r.body
}
