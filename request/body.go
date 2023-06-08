package request

import "github.com/ginger-core/errors"

func (r *request) ProcessBody(ref any) errors.Error {
	return nil
}

func (r *request) GetBody() any {
	return r.body
}
