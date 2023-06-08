package request

import "github.com/ginger-core/errors"

func (r *request) GetHeader(key string) string {
	return ""
}

func (r *request) ProcessHeaders(ref any) errors.Error {
	return nil
}

func (r *request) GetHeaders() any {
	return r.header
}
