package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"

	"github.com/mymmrac/mymm.gq/server/config"
	handlerPkg "github.com/mymmrac/mymm.gq/server/handler"
	"github.com/mymmrac/mymm.gq/server/logger"
)

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
	defer func() {
		if err = log.Close(); err != nil {
			exitWithError(err)
		}
	}()

	if cfg.Log.Level != config.LogLevelDebug {
		app.Configure(iris.WithoutStartupLog)
	}

	if cfg.CORSAllowAll {
		app.UseRouter(cors.AllowAll())
	}

	handler, err := handlerPkg.NewHandler(app, cfg, log)
	if err != nil {
		exitWithError(err)
	}

	handler.RegisterRoutes()

	idleConnectionsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		timeout := 10 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		_ = app.Shutdown(ctx)
		close(idleConnectionsClosed)
	})

	fmt.Println("Listening...")
	_ = app.Listen(":"+cfg.Port, iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	<-idleConnectionsClosed

	fmt.Println("Done")
}

func exitWithError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "FATAL: %s", err)
	os.Exit(1)
}
