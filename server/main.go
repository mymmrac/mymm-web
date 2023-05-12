package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/mymmrac/mymm-web/server/common"
	"github.com/mymmrac/mymm-web/server/config"
	handlerPkg "github.com/mymmrac/mymm-web/server/handler"
	"github.com/mymmrac/mymm-web/server/logger"
)

var configFile = flag.String("config", "", "Config file")

func main() {
	fmt.Println("Starting server...")

	app := iris.New()

	// ==== Config ====
	flag.Parse()
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		exitWithError(err)
	}

	if cfg.Infrastructure.HTTPServer.CORSAllowAll {
		app.UseRouter(cors.AllowAll())
	}
	// ==== Config End ====

	// ==== Logger ====
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
	// ==== Logger End ====

	// ==== MongoDB ====
	mongoCtx, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoCancel()
	mongoClient, err := mongo.Connect(mongoCtx, options.Client().
		ApplyURI("mongodb://"+cfg.Infrastructure.MongoDB.Host))
	defer func() {
		if err = mongoClient.Disconnect(mongoCtx); err != nil {
			exitWithError(err)
		}
	}()
	if err = mongoClient.Ping(mongoCtx, readpref.Primary()); err != nil {
		exitWithError(err)
	}
	// ==== MongoDB End ====

	// ==== Auth ====
	authCache := common.NewCachedAuth(cfg.Settings.AuthCacheExpiration.Duration)
	auth, err := common.SafeBasicAuthLoad(cfg.Settings.UsersFilename, basicauth.BCRYPT, authCache.Option)
	if err != nil {
		exitWithError(err)
	}
	// ==== Auth End ====

	handler, err := handlerPkg.NewHandler(app, auth, cfg, log, mongoClient)
	if err != nil {
		exitWithError(err)
	}

	handler.RegisterRoutes()

	idleConnectionsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		irisCtx, irisCancel := context.WithTimeout(context.Background(), cfg.Infrastructure.HTTPServer.Timeout.Duration)
		defer irisCancel()

		_ = app.Shutdown(irisCtx)
		close(idleConnectionsClosed)
	})

	fmt.Println("Listening...")
	_ = app.Listen("localhost:"+cfg.Infrastructure.HTTPServer.Port, iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed))
	<-idleConnectionsClosed

	fmt.Println("Done")
}

func exitWithError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "FATAL: %s", err)
	os.Exit(1)
}
