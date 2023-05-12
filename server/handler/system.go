package handler

import (
	"github.com/kataras/iris/v12/context"

	"github.com/mymmrac/mymm-web/server/common"
)

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
