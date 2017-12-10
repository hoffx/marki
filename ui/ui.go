package ui

import (
	"golang.org/x/text/language"
)

type UI struct {
	Langs []string
}

func (u *UI) MatchLang(langs []string) string {
	avail := make([]language.Tag, len(langs))
	for i := range avail {
		avail[i] = language.Make(langs[i])
	}
	matcher := language.NewMatcher(avail)
	desired := make([]language.Tag, len(u.Langs))
	for i := range desired {
		desired[i] = language.Make(u.Langs[i])
	}
	_, i, _ := matcher.Match(desired...)
	return langs[i]
}
