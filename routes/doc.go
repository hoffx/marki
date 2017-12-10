package routes

import (
	"io"
	"log"
	"strings"

	"github.com/hoffx/marki/storage"
	macaron "gopkg.in/macaron.v1"
)

func Doc(ctx *macaron.Context, log *log.Logger) {
	t, _ := storage.NewFSStorage("test-content").Root(storage.DefaultRevision)
	path := ctx.Params("*")
	// TODO: sanitize path?
	switch t.(type) {
	case storage.PathTopic:
		var err error
		t, err = t.(storage.PathTopic).Subtopic(path)
		if err != nil {
			ctx.WriteHeader(404)
			return
		}
	default:
		route := strings.Split(path, "/")
		for _, r := range route {
			for _, x := range t.Subtopics() {
				log.Println(x.ID())
				if x.ID() == r {
					t = x
					goto cont
				}
			}
			ctx.WriteHeader(404)
			return
		cont:
		}
	}

	d, _ := t.Data("de-DE")
	io.Copy(ctx.Resp, d)
}
