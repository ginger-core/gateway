package gateway

import (
	"github.com/ginger-core/errors"
)

type HandleFunc func(request Request) (any, errors.Error)

type Handler interface {
	Handle(request Request) (any, errors.Error)
}
