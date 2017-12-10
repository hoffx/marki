package ui

import (
	"html/template"

	"github.com/hoffx/marki/storage"
)

type TOC struct {
	Link      template.URL
	Name      string
	Subtopics []TOC
	Lang      string
}

func (u *UI) GenerateTOC(topic storage.Topic, basepath string) (toc *TOC) {
	toc = &TOC{
		Link: template.URL(basepath),
	}
	toc.Lang = u.MatchLang(topic.Locales())
	name, _ := topic.Name(toc.Lang)
	toc.Name = name
	st := topic.Subtopics()
	if len(st) >= 0 {
		toc.Subtopics = make([]TOC, len(st))
	}
	for i := range st {
		toc.Subtopics[i] = *u.GenerateTOC(st[i], basepath+"/"+topic.ID())
	}
	return
}
