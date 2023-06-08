package gateway

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Controller interface {
	WithLanguageBundle(bundle *i18n.Bundle) Controller
	GetLanguageBundle() *i18n.Bundle

	Responder
	Processor
}

type controller struct {
	Responder
	bundle *i18n.Bundle
}

func NewController(responder Responder) Controller {
	c := &controller{
		Responder: responder,
	}
	return c
}
