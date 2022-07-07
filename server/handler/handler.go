package handler

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mymmrac/mymm.gq/server/common"
	"github.com/mymmrac/mymm.gq/server/config"
	"github.com/mymmrac/mymm.gq/server/logger"
	"github.com/mymmrac/mymm.gq/server/model"
)

type Handler struct {
	app  *iris.Application
	auth context.Handler
	log  logger.Logger

	health    model.Health
	system    model.System
	bookmarks model.Bookmarks
}

func NewHandler(app *iris.Application, auth context.Handler, cfg config.Config, log logger.Logger,
	mongoClient *mongo.Client) (*Handler, error) {
	health, err := model.NewHealth()
	if err != nil {
		return nil, fmt.Errorf("failed to create health model: %w", err)
	}

	return &Handler{
		app:  app,
		auth: auth,
		log:  log,

		health:    health,
		system:    model.NewSystem(cfg),
		bookmarks: model.NewBookmarks(log, mongoClient),
	}, nil
}

func (h *Handler) RegisterRoutes() {
	h.handleErrors()

	h.app.Get("/", h.healthHandler)

	systemAPI := h.app.Party("/system", h.auth)
	systemAPI.Get("/", h.systemAllHandler)
	systemAPI.Get("/cpu", h.cpuHandler)
	systemAPI.Get("/load", h.loadHandler)
	systemAPI.Get("/ram", h.ramHandler)
	systemAPI.Get("/swap", h.swapHandler)
	systemAPI.Get("/uptime", h.uptimeHandler)
	systemAPI.Get("/disk", h.diskHandler)

	bookmarksAPI := h.app.Party("/bookmarks")
	bookmarksAPI.Get("/", h.bookmarksHandler)
	bookmarksAPI.Post("/", h.auth, h.bookmarksAddHandler)
	bookmarksAPI.Put("/", h.auth, h.bookmarksUpdateHandler)
	bookmarksAPI.Delete("/", h.auth, h.bookmarksDeleteHandler)
}

func (h *Handler) handleErrors() {
	h.app.OnErrorCode(iris.StatusUnauthorized, func(ctx *context.Context) {
		common.ReturnErrorTextWithStatus(ctx, iris.StatusUnauthorized, "401 Unauthorized")
	})
	h.app.OnErrorCode(iris.StatusNotFound, func(ctx *context.Context) {
		common.ReturnErrorTextWithStatus(ctx, iris.StatusNotFound, "404 Not Found")
	})
	h.app.OnErrorCode(iris.StatusInternalServerError, func(ctx *context.Context) {
		common.ReturnErrorTextWithStatus(ctx, iris.StatusInternalServerError, "500 Internal Server Error")
	})
}
