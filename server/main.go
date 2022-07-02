package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/mymmrac/mymm.gq/server/common"
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

	mongoCtx, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoCancel()
	client, err := mongo.Connect(mongoCtx, options.Client().
		SetRegistry(common.MongoRegistry).
		ApplyURI("mongodb://"+cfg.MongoDBHost))
	defer func() {
		if err = client.Disconnect(mongoCtx); err != nil {
			exitWithError(err)
		}
	}()
	if err = client.Ping(mongoCtx, readpref.Primary()); err != nil {
		exitWithError(err)
	}

	handler, err := handlerPkg.NewHandler(app, cfg, log, client)
	if err != nil {
		exitWithError(err)
	}

	handler.RegisterRoutes()

	idleConnectionsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		timeout := 10 * time.Second
		irisCtx, irisCancel := context.WithTimeout(context.Background(), timeout)
		defer irisCancel()

		_ = app.Shutdown(irisCtx)
		close(idleConnectionsClosed)
	})

	fmt.Println("Listening...")
	_ = app.Listen("localhost:"+cfg.Port, iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	<-idleConnectionsClosed

	fmt.Println("Done")
}

func exitWithError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "FATAL: %s", err)
	os.Exit(1)
}
