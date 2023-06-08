package gateway

import (
	"github.com/ginger-core/errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func TranslateError(language Language, err errors.Error) errors.Error {
	id := err.GetId()
	if id != "" {
		msg, _ := language.Localize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    id,
				Other: err.GetMessage(),
				One:   err.GetMessageOne(),
			},
			TemplateData: err.GetProperties(),
			PluralCount:  err.GetPluralCount(),
		})
		if msg == "" {
			msg = err.GetMessage()
		}
		err = err.Clone().WithMessage(msg)
	}

	return err
}
