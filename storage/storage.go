package storage

import (
	"errors"
	"io"
)

type Revision string

const DefaultRevision Revision = "master"

var (
	ErrNoSuchRevision = errors.New("no such revision")
	ErrNoSuchLocale   = errors.New("no such locale")
	ErrNoSuchSubtopic = errors.New("no such subtopic")
)

type Storage interface {
	Root(revision Revision) (Topic, error)
	Sync()
	GetRevisions() []Revision
}

// Topic depicts an article and its subtopics
type Topic interface {
	// ID returns an un-localized name for the article
	ID() string

	// Locales returns all available locales
	Locales() []string

	// Name returns the localized Name
	Name(locale string) (string, error) // non-unique name

	// Data returns a rewinded io.Reader delivering the markdown code
	Data(locale string) (io.Reader, error) // link

	// Subtopics returns the subtopics of the topic
	Subtopics() []Topic // subtopics
}

// PathTopic is a Topic that can be queried by a path
type PathTopic interface {
	Topic

	// Subtopic returns a subtopic in a given path
	Subtopic(path string) (Topic, error)
}
