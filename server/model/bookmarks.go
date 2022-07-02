package model

import "github.com/google/uuid"

type BookmarkAddRequest struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	IconURL  string `json:"iconURL"`
	Category string `json:"category"`
}

type Bookmark struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	IconURL  string    `json:"iconURL"`
	Category string    `json:"category"`
}

type Bookmarks interface {
	Bookmarks() ([]Bookmark, error)
	Add(request BookmarkAddRequest) error
	Delete(id uuid.UUID) error
	Update(bookmark Bookmark) error
}

type BookmarksImpl struct{}

func (b *BookmarksImpl) Bookmarks() ([]Bookmark, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookmarksImpl) Add(request BookmarkAddRequest) error {
	//TODO implement me
	panic("implement me")
}

func (b *BookmarksImpl) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (b *BookmarksImpl) Update(bookmark Bookmark) error {
	//TODO implement me
	panic("implement me")
}
