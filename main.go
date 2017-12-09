package main

import (
	"os"

	"github.com/hoffx/marki/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Marki"
	app.Usage = "Markdown rendering server"
	app.Author = "Hoff Industires"
	app.Commands = []cli.Command{cmd.Web}
	app.Run(os.Args)
}
