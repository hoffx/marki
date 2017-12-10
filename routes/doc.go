package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hoffx/marki/storage"
	"github.com/hoffx/marki/ui"
	macaron "gopkg.in/macaron.v1"
)

func Doc(ctx *macaron.Context, log *log.Logger, ui *ui.UI) {
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
	toc := ui.GenerateTOC(t, ctx.Req.URL.Path)
	fmt.Printf("%+v\n", toc)
	ctx.Data["Toc"] = toc
	ctx.Data["Name"] = toc.Name
	data, _ := t.Data(toc.Lang)
	d, _ := ioutil.ReadAll(data)
	ctx.Data["Data"] = string(d)
	ctx.HTML(200, "doc")
}
