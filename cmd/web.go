package cmd

import (
	"github.com/go-macaron/i18n"
	"github.com/hoffx/marki/middleware"
	"github.com/hoffx/marki/routes"
	"github.com/urfave/cli"
	macaron "gopkg.in/macaron.v1"
)

var Web = cli.Command{
	Name:   "web",
	Usage:  "Start webserver",
	Action: runWeb,
}

func runWeb(ctx *cli.Context) {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("static", macaron.StaticOptions{
		SkipLogging: true,
	}))
	m.Use(i18n.I18n(i18n.Options{
		Directory: "locales",
		Langs:     []string{"de-DE", "en-US"},
		Names:     []string{"Deutsch", "Englisch"},
	}))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "templates",
	}))
	m.Use(middleware.UI)
	m.Get("/", routes.Home)
	m.Get("/*", routes.Doc)
	m.Run()
}
