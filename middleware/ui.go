package middleware

import (
	"github.com/go-macaron/i18n"
	"github.com/hoffx/marki/ui"
	macaron "gopkg.in/macaron.v1"
)

func UI(ctx *macaron.Context, locale i18n.Locale) {
	ctx.Map(&ui.UI{
		Langs: []string{locale.Lang},
	})
}
