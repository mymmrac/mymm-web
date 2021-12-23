package main

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

const cpuReadDuration = 1 * time.Second

type cpuLoad struct {
	Cores []float64 `json:"cores"`
}

func cpuHandler(ctx *context.Context) {
	load, err := cpu.Percent(cpuReadDuration, true)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading CPU").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(cpuLoad{
		Cores: load,
	})
}

type ramUsage struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func ramHandler(ctx *context.Context) {
	ram, err := mem.VirtualMemory()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading RAM").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(ramUsage{
		Total:       ram.Total,
		Free:        ram.Free,
		Used:        ram.Used,
		UsedPercent: ram.UsedPercent,
	})
}

type swapUsage struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func swapHandler(ctx *context.Context) {
	swap, err := mem.SwapMemory()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading RAM").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(swapUsage{
		Total:       swap.Total,
		Free:        swap.Free,
		Used:        swap.Used,
		UsedPercent: swap.UsedPercent,
	})
}

type diskUsage struct {
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func diskHandler(ctx *context.Context) {
	diskUsg, err := disk.Usage("/")
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading RAM").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(diskUsage{
		Path:        diskUsg.Path,
		Total:       diskUsg.Total,
		Free:        diskUsg.Free,
		Used:        diskUsg.Used,
		UsedPercent: diskUsg.UsedPercent,
	})
}
