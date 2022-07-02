package main

import (
	"errors"
	"math/rand"
	"runtime/debug"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func registerRoutes(app *iris.Application) {
	app.Get("/", healthRouter)

	systemAPI := app.Party("/system")

	systemAPI.Get("/cpu", cpuHandler)
	systemAPI.Get("/ram", ramHandler)
	systemAPI.Get("/swap", swapHandler)
	systemAPI.Get("/disk", diskHandler)
	systemAPI.Get("/uptime", uptimeHandler)
	systemAPI.Get("/load", loadHandler)
}

type healthStats struct {
	Running   bool      `json:"running"`
	Random    int       `json:"random"`
	Time      string    `json:"time"`
	BuildInfo buildInfo `json:"buildInfo"`
}

type buildInfo struct {
	Module       string            `json:"module"`
	Version      string            `json:"version"`
	Dependencies int               `json:"dependencies"`
	Settings     map[string]string `json:"settings"`
}

func healthRouter(ctx *context.Context) {
	build, ok := debug.ReadBuildInfo()
	if !ok {
		ctx.StopWithError(iris.StatusInternalServerError, errors.New("build info is unavailable"))
		return
	}

	settings := make(map[string]string)
	for _, setting := range build.Settings {
		settings[setting.Key] = setting.Value
	}

	_, _ = ctx.JSON(healthStats{
		Running: true,
		Random:  rand.Int(),
		Time:    time.Now().Local().String(),
		BuildInfo: buildInfo{
			Module:       build.Main.Path,
			Version:      build.Main.Version,
			Dependencies: len(build.Deps),
			Settings:     settings,
		},
	})
}
