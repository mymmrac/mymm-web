package model

import (
	"errors"
	"math/rand"
	"runtime/debug"
	"time"
)

type HealthStats struct {
	Running   bool      `json:"running"`
	Random    int       `json:"random"`
	Time      string    `json:"time"`
	BuildInfo BuildInfo `json:"buildInfo"`
}

type BuildInfo struct {
	Module       string            `json:"module"`
	Version      string            `json:"version"`
	Dependencies int               `json:"dependencies"`
	Settings     map[string]string `json:"settings"`
}

type Health interface {
	Health() (*HealthStats, error)
}

type HealthImpl struct {
	build BuildInfo
}

func NewHealth() (*HealthImpl, error) {
	build, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("failed to get build info")
	}

	settings := make(map[string]string)
	for _, setting := range build.Settings {
		settings[setting.Key] = setting.Value
	}

	return &HealthImpl{
		build: BuildInfo{
			Module:       build.Main.Path,
			Version:      build.Main.Version,
			Dependencies: len(build.Deps),
			Settings:     settings,
		},
	}, nil
}

func (h *HealthImpl) Health() (*HealthStats, error) {
	return &HealthStats{
		Running:   true,
		Random:    rand.Int(),
		Time:      time.Now().Local().String(),
		BuildInfo: h.build,
	}, nil
}
