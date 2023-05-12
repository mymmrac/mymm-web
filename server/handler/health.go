package handler

import (
	"github.com/kataras/iris/v12/context"

	"github.com/mymmrac/mymm-web/server/common"
)

func (h *Handler) healthHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.health.Health)
}
