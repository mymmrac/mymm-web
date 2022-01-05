package main

import (
	stdContext "context"
	"flag"
	"math/rand"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

const addr = "127.0.0.1:8080"

const logTimeFormat = "02.01.2006 15:04"

type healthStats struct {
	Running bool `json:"running"`
	Random  int  `json:"random"`
}

var localRun = flag.Bool("local", false, "Local run")

func main() {
	app := iris.New()
	app.Logger().TimeFormat = logTimeFormat

	flag.Parse()
	if *localRun {
		crs := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
		})
		app.UseRouter(crs)
	}

	app.Get("/", func(ctx *context.Context) {
		_, _ = ctx.JSON(healthStats{
			Running: true,
			Random:  rand.Int(),
		})
	})

	systemAPI := app.Party("/system")

	systemAPI.Get("/cpu", cpuHandler)
	systemAPI.Get("/ram", ramHandler)
	systemAPI.Get("/swap", swapHandler)
	systemAPI.Get("/disk", diskHandler)
	systemAPI.Get("/uptime", uptimeHandler)
	systemAPI.Get("/load", loadHandler)

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
