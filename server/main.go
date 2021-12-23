package main

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

const addr = "127.0.0.1:8080"

const logTimeFormat = "02.01.2006 15:04"

const cpuReadDuration = 1 * time.Second

func main() {
	app := iris.New()
	app.Logger().TimeFormat = logTimeFormat
	systemAPI := app.Party("/system")

	systemAPI.Get("/cpu", func(ctx *context.Context) {
		load, err := cpu.Percent(cpuReadDuration, true)
		if err != nil {
			ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
				Title("Reading CPU").DetailErr(err))
			return
		}

		_, _ = ctx.JSON(load)
	})

	systemAPI.Get("/ram", func(ctx *context.Context) {
		ram, err := mem.VirtualMemory()
		if err != nil {
			ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
				Title("Reading RAM").DetailErr(err))
			return
		}

		_, _ = ctx.JSON(ram)
	})

	_ = app.Listen(addr)
}
