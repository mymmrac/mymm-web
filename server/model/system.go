package model

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

type CpuStats struct {
	Cores []float64 `json:"cores"`
}

type LoadStats struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

type RamStats struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type SwapStats struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type UptimeStats struct {
	BootTime uint64 `json:"bootTime"`
	Uptime   uint64 `json:"uptime"`
}

type DiskStats struct {
	Path        string  `json:"path"`
	FSType      string  `json:"fsType"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type System interface {
	CPU() (*CpuStats, error)
	Load() (*LoadStats, error)
	RAM() (*RamStats, error)
	Swap() (*SwapStats, error)
	Uptime() (*UptimeStats, error)
	Disk() (*DiskStats, error)
}

type SystemImpl struct{}

func NewSystem() *SystemImpl {
	return &SystemImpl{}
}

const cpuReadDuration = 1 * time.Second

func (s *SystemImpl) CPU() (*CpuStats, error) {
	cores, err := cpu.Percent(cpuReadDuration, true)
	if err != nil {
		return nil, fmt.Errorf("failed to read cpu data: %w", err)
	}

	return &CpuStats{
		Cores: cores,
	}, nil
}

func (s *SystemImpl) Load() (*LoadStats, error) {
	avgLoad, err := load.Avg()
	if err != nil {
		return nil, fmt.Errorf("failed to read load data: %w", err)
	}

	return &LoadStats{
		Load1:  avgLoad.Load1,
		Load5:  avgLoad.Load5,
		Load15: avgLoad.Load15,
	}, nil
}

func (s *SystemImpl) RAM() (*RamStats, error) {
	ram, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to read ram data: %w", err)
	}

	return &RamStats{
		Total:       ram.Total,
		Free:        ram.Free,
		Used:        ram.Used,
		UsedPercent: ram.UsedPercent,
	}, nil
}

func (s *SystemImpl) Swap() (*SwapStats, error) {
	swap, err := mem.SwapMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to read swap data: %w", err)
	}

	return &SwapStats{
		Total:       swap.Total,
		Free:        swap.Free,
		Used:        swap.Used,
		UsedPercent: swap.UsedPercent,
	}, nil
}

func (s *SystemImpl) Uptime() (*UptimeStats, error) {
	uptime, err := host.Uptime()
	if err != nil {
		return nil, fmt.Errorf("failed to read uptime data: %w", err)
	}

	bootTime, err := host.BootTime()
	if err != nil {
		return nil, fmt.Errorf("failed to read boot time data: %w", err)
	}

	return &UptimeStats{
		BootTime: bootTime,
		Uptime:   uptime,
	}, nil
}

func (s *SystemImpl) Disk() (*DiskStats, error) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("failed to read disk data: %w", err)
	}

	return &DiskStats{
		Path:        diskUsage.Path,
		FSType:      diskUsage.Fstype,
		Total:       diskUsage.Total,
		Free:        diskUsage.Free,
		Used:        diskUsage.Used,
		UsedPercent: diskUsage.UsedPercent,
	}, nil
}
