package gateway

import (
	"github.com/ginger-core/errors"
)

type Binding interface {
	// Bind tries to bind the request to given ref
	Bind(request Request, ref any) errors.Error
}
