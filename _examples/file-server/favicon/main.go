package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()

	// This will serve the ./static/favicons/favicon.ico to: localhost:8080/favicon.ico
	app.Favicon("./static/favicons/favicon.ico.ico")

	// app.Favicon("./static/favicons/favicon.ico.ico", "/favicon_16_16.ico")
	// This will serve the ./static/favicons/favicon.ico.ico to: localhost:8080/favicon_16_16.ico

	app.Get("/", func(ctx context.Context) {
		ctx.HTML(`<a href="/favicon.ico"> press here to see the favicon.ico</a>.
		 At some browsers like chrome, it should be visible at the top-left side of the browser's window,
		 because some browsers make requests to the /favicon.ico automatically,
		  so ion serves your favicon in that path too (you can change it).`)
	}) // if favicon doesn't show to you, try to clear your browser's cache.

	app.Run(ion.Addr(":8080"))
}
