package main

import (
	"github.com/shirou/gopsutil/v3/load"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

const cpuReadDuration = 1 * time.Second

type cpuStats struct {
	Cores []float64 `json:"cores"`
}

func cpuHandler(ctx *context.Context) {
	cores, err := cpu.Percent(cpuReadDuration, true)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading CPU").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(cpuStats{
		Cores: cores,
	})
}

type ramStats struct {
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

	_, _ = ctx.JSON(ramStats{
		Total:       ram.Total,
		Free:        ram.Free,
		Used:        ram.Used,
		UsedPercent: ram.UsedPercent,
	})
}

type swapStats struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func swapHandler(ctx *context.Context) {
	swap, err := mem.SwapMemory()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading swap").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(swapStats{
		Total:       swap.Total,
		Free:        swap.Free,
		Used:        swap.Used,
		UsedPercent: swap.UsedPercent,
	})
}

type diskStats struct {
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func diskHandler(ctx *context.Context) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading disk").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(diskStats{
		Path:        diskUsage.Path,
		Total:       diskUsage.Total,
		Free:        diskUsage.Free,
		Used:        diskUsage.Used,
		UsedPercent: diskUsage.UsedPercent,
	})
}

type uptimeStats struct {
	BootTime uint64 `json:"bootTime"`
	Uptime   uint64 `json:"uptime"`
}

func uptimeHandler(ctx *context.Context) {
	uptime, err := host.Uptime()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading uptime").DetailErr(err))
		return
	}

	bootTime, err := host.BootTime()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading boot time").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(uptimeStats{
		BootTime: bootTime,
		Uptime:   uptime,
	})
}

type loadStats struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

func loadHandler(ctx *context.Context) {
	avgLoad, err := load.Avg()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Reading disk").DetailErr(err))
		return
	}

	_, _ = ctx.JSON(loadStats{
		Load1:  avgLoad.Load1,
		Load5:  avgLoad.Load5,
		Load15: avgLoad.Load15,
	})
}
