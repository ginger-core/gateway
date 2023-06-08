package gateway

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Language interface {
	Localize(lc *i18n.LocalizeConfig) (string, error)
}

type language struct {
	bundle *i18n.Bundle
	*i18n.Localizer
}

func NewLanguage(bundle *i18n.Bundle, langs ...string) Language {
	l := &language{
		bundle:    bundle,
		Localizer: i18n.NewLocalizer(bundle, langs...),
	}
	return l
}

/**/

func (c *controller) WithLanguageBundle(bundle *i18n.Bundle) Controller {
	c.bundle = bundle
	return c
}

func (c *controller) GetLanguageBundle() *i18n.Bundle {
	return c.bundle
}
