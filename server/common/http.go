package common

import (
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

func ReturnJSON[T any](ctx *context.Context, f func() (T, error)) {
	data, err := f()
	if err != nil {
		ReturnError(ctx, err)
		return
	}

	if _, err = ctx.JSON(data); err != nil {
		ReturnError(ctx, err)
	}
}

type responseError struct {
	Error string `json:"error"`
}

func ReturnError(ctx *context.Context, err error) {
	status := iris.StatusInternalServerError
	if errors.Is(err, ErrNotFound) {
		status = iris.StatusNotFound
	}

	if sendErr := ctx.StopWithJSON(status, responseError{Error: err.Error()}); sendErr != nil {
		ctx.StopWithError(iris.StatusInternalServerError, fmt.Errorf(
			"failed to send error: %s, original error: %w", sendErr, err))
	}
}

func ReturnErrorText(ctx *context.Context, errText string) {
	if sendErr := ctx.StopWithJSON(iris.StatusInternalServerError, responseError{Error: errText}); sendErr != nil {
		ctx.StopWithError(iris.StatusInternalServerError, fmt.Errorf(
			"failed to send error: %w, original error: %s", sendErr, errText))
	}
}

func ReturnErrorWithStatus(ctx *context.Context, status int, err error) {
	if sendErr := ctx.StopWithJSON(status, responseError{Error: err.Error()}); sendErr != nil {
		ctx.StopWithError(iris.StatusInternalServerError, fmt.Errorf(
			"failed to send error: %s, original error: %w", sendErr, err))
	}
}

func ReturnErrorTextWithStatus(ctx *context.Context, status int, errText string) {
	if sendErr := ctx.StopWithJSON(status, responseError{Error: errText}); sendErr != nil {
		ctx.StopWithError(iris.StatusInternalServerError, fmt.Errorf(
			"failed to send error: %w, original error: %s", sendErr, errText))
	}
}

func SafeBasicAuthLoad(filename string, userOptions ...basicauth.UserAuthOption) (handler context.Handler, err error) {
	defer func() {
		result := recover()
		if result != nil {
			if panicErr, ok := result.(error); ok {
				err = fmt.Errorf("failed to create auth: %w", panicErr)
			} else {
				panic(result)
			}
		}
	}()

	handler = basicauth.Load(filename, userOptions...)
	return handler, err
}
