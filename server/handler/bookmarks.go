package handler

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

	"github.com/mymmrac/mymm.gq/server/common"
	"github.com/mymmrac/mymm.gq/server/model"
)

func (h *Handler) bookmarksHandler(ctx *context.Context) {
	common.ReturnJSON(ctx, h.bookmarks.Bookmarks)
}

func (h *Handler) bookmarksAddHandler(ctx *context.Context) {
	var request model.BookmarkAddRequest
	if err := ctx.ReadJSON(&request); err != nil {
		common.ReturnErrorWithStatus(ctx, iris.StatusBadRequest, err)
		return
	}

	if err := h.bookmarks.Add(request); err != nil {
		common.ReturnError(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusOK)
}

func (h *Handler) bookmarksUpdateHandler(ctx *context.Context) {
	var bookmark model.Bookmark
	if err := ctx.ReadJSON(&bookmark); err != nil {
		common.ReturnErrorWithStatus(ctx, iris.StatusBadRequest, err)
		return
	}

	if bookmark.ID == uuid.Nil {
		common.ReturnErrorTextWithStatus(ctx, iris.StatusBadRequest, "bookmark ID is empty")
		return
	}

	if err := h.bookmarks.Update(bookmark); err != nil {
		common.ReturnError(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusOK)
}

func (h *Handler) bookmarksDeleteHandler(ctx *context.Context) {
	var bookmarkID uuid.UUID
	if err := ctx.ReadJSON(&bookmarkID); err != nil {
		common.ReturnErrorWithStatus(ctx, iris.StatusBadRequest, err)
		return
	}

	if bookmarkID == uuid.Nil {
		common.ReturnErrorTextWithStatus(ctx, iris.StatusBadRequest, "bookmark ID is empty")
		return
	}

	if err := h.bookmarks.Delete(bookmarkID); err != nil {
		common.ReturnError(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusOK)
}
