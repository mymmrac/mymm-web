package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func returnJSON[T any](ctx *context.Context, f func() (T, error)) {
	data, err := f()
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	if _, err = ctx.JSON(data); err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
	}
}
