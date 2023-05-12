package model

import (
	"fmt"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/mymmrac/mymm-web/server/common"
	"github.com/mymmrac/mymm-web/server/config"
)

type CPUStats struct {
	Cores []float64 `json:"cores"`
}

type LoadStats struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

type RAMStats struct {
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

type SystemStats struct {
	CPU    *CPUStats    `json:"cpu"`
	Load   *LoadStats   `json:"load"`
	RAM    *RAMStats    `json:"ram"`
	Swap   *SwapStats   `json:"swap"`
	Uptime *UptimeStats `json:"uptime"`
	Disk   *DiskStats   `json:"disk"`
}

type System interface {
	All() (*SystemStats, error)

	CPU() (*CPUStats, error)
	Load() (*LoadStats, error)
	RAM() (*RAMStats, error)
	Swap() (*SwapStats, error)
	Uptime() (*UptimeStats, error)
	Disk() (*DiskStats, error)
}

type SystemImpl struct {
	cpuReadDuration time.Duration
}

func NewSystem(cfg config.Config) *SystemImpl {
	return &SystemImpl{
		cpuReadDuration: cfg.Settings.CPUReadDuration.Duration,
	}
}

func (s *SystemImpl) All() (*SystemStats, error) {
	var stats SystemStats
	var cpuErr, loadErr, ramErr, swapErr, uptimeErr, diskErr error

	wg := sync.WaitGroup{}
	wg.Add(6)

	go func() { stats.CPU, cpuErr = s.CPU(); wg.Done() }()
	go func() { stats.Load, loadErr = s.Load(); wg.Done() }()
	go func() { stats.RAM, ramErr = s.RAM(); wg.Done() }()
	go func() { stats.Swap, swapErr = s.Swap(); wg.Done() }()
	go func() { stats.Uptime, uptimeErr = s.Uptime(); wg.Done() }()
	go func() { stats.Disk, diskErr = s.Disk(); wg.Done() }()

	wg.Wait()
	if err := common.FirstError(cpuErr, loadErr, ramErr, swapErr, uptimeErr, diskErr); err != nil {
		return nil, err
	}

	return &stats, nil
}

func (s *SystemImpl) CPU() (*CPUStats, error) {
	cores, err := cpu.Percent(s.cpuReadDuration, true)
	if err != nil {
		return nil, fmt.Errorf("failed to read cpu data: %w", err)
	}

	return &CPUStats{
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

func (s *SystemImpl) RAM() (*RAMStats, error) {
	ram, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to read ram data: %w", err)
	}

	return &RAMStats{
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
