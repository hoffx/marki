package storage

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type fsStorage struct {
	path string
}

func (s *fsStorage) GetRevisions() []Revision {
	return []Revision{DefaultRevision}
}

func (s *fsStorage) Root(r Revision) (Topic, error) {
	if r != DefaultRevision {
		return nil, ErrNoSuchRevision
	}
	return &fsTopic{
		path: path.Clean(s.path),
	}, nil
}

func (s *fsStorage) Sync() {

}

type fsTopic struct {
	path string
}

func (t *fsTopic) ID() string {
	return path.Base(t.path)
}

func (t *fsTopic) Locales() (ls []string) {
	f, err := ioutil.ReadDir(t.path)
	if err != nil {
		panic(err)
	}
	ls = make([]string, 0)
	for _, l := range f {
		if !l.IsDir() && strings.HasSuffix(l.Name(), ".md") {
			ls = append(ls, strings.TrimSuffix(l.Name(), ".md"))
		}
	}
	return
}

func (t *fsTopic) Name(locale string) (string, error) {
	return t.ID(), nil
}

func (t *fsTopic) Data(locale string) (io.Reader, error) {
	r, err := os.Open(path.Join(t.path, locale+".md"))
	if err == os.ErrNotExist {
		return nil, ErrNoSuchLocale
	}
	if err != nil {
		panic(err)
	}
	return r, nil
}

func (t *fsTopic) Subtopics() (ts []Topic) {
	f, err := ioutil.ReadDir(t.path)
	if err != nil {
		panic(err)
	}
	ts = make([]Topic, 0)
	for _, l := range f {
		if l.IsDir() {
			ts = append(ts, &fsTopic{
				path: path.Join(t.path, l.Name()),
			})
		}
	}
	return
}

func NewFSStorage(path string) Storage {
	return &fsStorage{
		path: path,
	}
}
