package gateway

import (
	"time"

	"github.com/ginger-core/errors"
)

type Server interface {
	NewResponder() Responder

	SetController(controller Controller)
	GetController() Controller

	NewRouterGroup(path string) RouterGroup

	Run() errors.Error
	Shutdown(timeout time.Duration) errors.Error
}
