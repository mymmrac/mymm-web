package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

	"github.com/mymmrac/mymm.gq/server/common"
	"github.com/mymmrac/mymm.gq/server/config"
	"github.com/mymmrac/mymm.gq/server/logger"
	"github.com/mymmrac/mymm.gq/server/model"
)

type Handler struct {
	app *iris.Application
	cfg config.Config
	log logger.Logger

	health model.Health
	system model.System
}

func NewHandler(app *iris.Application, cfg config.Config, log logger.Logger) (*Handler, error) {
	health, err := model.NewHealth()
	if err != nil {
		return nil, fmt.Errorf("failed to create health model: %w", err)
	}

	return &Handler{
		app: app,
		cfg: cfg,
		log: log,

		health: health,
		system: model.NewSystem(cfg),
	}, nil
}

func (h *Handler) RegisterRoutes() {
	h.app.Get("/", h.healthHandler)

	systemAPI := h.app.Party("/system")

	systemAPI.Get("/", h.systemAllHandler)

	systemAPI.Get("/cpu", h.cpuHandler)
	systemAPI.Get("/load", h.loadHandler)
	systemAPI.Get("/ram", h.ramHandler)
	systemAPI.Get("/swap", h.swapHandler)
	systemAPI.Get("/uptime", h.uptimeHandler)
	systemAPI.Get("/disk", h.diskHandler)
}

func (h *Handler) healthHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.health.Health)
}

func (h *Handler) systemAllHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.system.All)
}

func (h *Handler) cpuHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.system.CPU)
}

func (h *Handler) loadHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.system.Load)
}

func (h *Handler) ramHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.system.RAM)
}

func (h *Handler) swapHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.system.Swap)
}

func (h *Handler) uptimeHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.system.Uptime)
}

func (h *Handler) diskHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.system.Disk)
}
