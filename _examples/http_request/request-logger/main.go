package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"github.com/get-ion/ion/middleware/logger"
)

func main() {
	app := ion.New()

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	app.Use(customLogger)

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("hello")
	})

	app.Get("/1", func(ctx context.Context) {
		ctx.Writef("hello")
	})

	app.Get("/2", func(ctx context.Context) {
		ctx.Writef("hello")
	})

	// log http errors should be done manually
	errorLogger := logger.New()

	app.OnErrorCode(ion.StatusNotFound, func(ctx context.Context) {
		errorLogger(ctx)
		ctx.Writef("My Custom 404 error page ")
	})

	// http://localhost:8080
	// http://localhost:8080/1
	// http://localhost:8080/2
	// http://lcoalhost:8080/notfoundhere
	// see the output on the console.
	app.Run(ion.Addr(":8080"))

}
