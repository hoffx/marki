package routes

import macaron "gopkg.in/macaron.v1"

func Doc(ctx *macaron.Context) {
	// TODO: parse route
	ctx.HTML(200, "doc")
}
