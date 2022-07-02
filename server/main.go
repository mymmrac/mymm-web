package main

import (
	stdContext "context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

	"github.com/mymmrac/mymm.gq/server/config"
	"github.com/mymmrac/mymm.gq/server/logger"
)

type healthStats struct {
	Running bool `json:"running"`
	Random  int  `json:"random"`
}

var configFile = flag.String("config", "", "Config file")

func main() {
	fmt.Println("Starting server...")

	flag.Parse()
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		exitWithError(err)
	}

	app := iris.New()

	log := logger.NewLog(app.Logger())
	if err = cfg.ConfigureLogger(log); err != nil {
		exitWithError(err)
	}

	if cfg.CORSAllowAll {
		app.UseRouter(cors.AllowAll())
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

	_ = app.Listen(":"+cfg.Port, iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	<-idleConnectionsClosed
}

func exitWithError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "FATAL: %s", err)
	os.Exit(1)
}
