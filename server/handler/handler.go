package handler

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"github.com/mymmrac/mymm.gq/server/config"
	"github.com/mymmrac/mymm.gq/server/logger"
	"github.com/mymmrac/mymm.gq/server/model"
)

type Handler struct {
	app *iris.Application
	cfg config.Config
	log logger.Logger

	health    model.Health
	system    model.System
	bookmarks model.Bookmarks
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

	bookmarksAPI := h.app.Party("/bookmarks")
	bookmarksAPI.Get("/", h.bookmarksHandler)
	bookmarksAPI.Post("/", h.bookmarksAddHandler)
	bookmarksAPI.Put("/", h.bookmarksUpdateHandler)
	bookmarksAPI.Delete("/", h.bookmarksDeleteHandler)
}
