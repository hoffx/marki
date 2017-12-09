package cmd

import (
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
	m.Use(macaron.Static("public", macaron.StaticOptions{
		SkipLogging: true,
	}))
	//m.Use(i18n.I18n())
	m.Get("/", routes.Home)
	m.Get("/*", routes.Doc)
	m.Run()
}
