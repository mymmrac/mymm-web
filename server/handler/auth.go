package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func (h *Handler) loginCheckHandler(ctx *context.Context) {
	ctx.StatusCode(iris.StatusAccepted)
}
