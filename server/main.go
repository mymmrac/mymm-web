package main

import (
	stdContext "context"
	"time"

	"github.com/kataras/iris/v12"
)

const addr = "127.0.0.1:8080"

const logTimeFormat = "02.01.2006 15:04"

func main() {
	app := iris.New()
	app.Logger().TimeFormat = logTimeFormat

	systemAPI := app.Party("/system")

	systemAPI.Get("/cpu", cpuHandler)
	systemAPI.Get("/ram", ramHandler)
	systemAPI.Get("/swap", swapHandler)
	systemAPI.Get("/disk", diskHandler)

	idleConnectionsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		timeout := 10 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()

		_ = app.Shutdown(ctx)
		close(idleConnectionsClosed)
	})

	_ = app.Listen(addr, iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	<-idleConnectionsClosed
}
