package storage

import "io"

type Revision string

type Storage interface {
	Root(revision Revision) Topic
	Sync()
	GetRevisions()
}

// Topic depicts an article and its subtopics
type Topic interface {
	// ID returns an un-localized name for the article
	ID() string

	// Locales returns all available locales
	Locales() []string

	// Name returns the localized Name
	Name(locale string) string // non-unique name

	// Data returns a rewinded io.Reader delivering the markdown code
	Data(locale string) io.Reader // link

	// Subtopics returns the subtopics of the topic
	Subtopics() []Topic // subtopics
}
