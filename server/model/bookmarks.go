package model

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mymmrac/mymm.gq/server/common"
)

type BookmarkAddRequest struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	IconURL  string `json:"iconUrl"`
	Category string `json:"category"`
}

type Bookmark struct {
	ID       uuid.UUID `json:"id"       bson:"_id"`
	Name     string    `json:"name"     bson:"name"`
	URL      string    `json:"url"      bson:"url"`
	IconURL  string    `json:"iconUrl"  bson:"iconUrl"`
	Category string    `json:"category" bson:"category"`
}

type Bookmarks interface {
	Bookmarks() ([]Bookmark, error)
	Add(request BookmarkAddRequest) error
	Delete(id uuid.UUID) error
	Update(bookmark Bookmark) error
}

type BookmarksImpl struct {
	collection *mongo.Collection
}

func NewBookmarks(client *mongo.Client) *BookmarksImpl {
	collection := client.Database("mymm").Collection("bookmarks")
	return &BookmarksImpl{
		collection: collection,
	}
}

func (b *BookmarksImpl) Bookmarks() ([]Bookmark, error) {
	ctx := context.TODO()
	cursor, err := b.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to get bookmarks: %w", err)
	}

	var bookmarks []Bookmark
	if err = cursor.All(ctx, &bookmarks); err != nil {
		return nil, fmt.Errorf("failed to decode bookmarks: %w", err)
	}

	return bookmarks, nil
}

func (b *BookmarksImpl) Add(request BookmarkAddRequest) error {
	bookmark := Bookmark{
		ID:       uuid.New(),
		Name:     request.Name,
		URL:      request.URL,
		IconURL:  request.IconURL,
		Category: request.Category,
	}

	_, err := b.collection.InsertOne(context.TODO(), bookmark)
	if err != nil {
		return fmt.Errorf("failed to add bookmark: %w", err)
	}

	return nil
}

func (b *BookmarksImpl) Delete(id uuid.UUID) error {
	result, err := b.collection.DeleteOne(context.TODO(),
		bson.D{{Key: "_id", Value: id}},
		options.Delete().SetHint(bson.D{{"_id", 1}}),
	)
	if err != nil {
		return fmt.Errorf("failed to delete bookmark: %w", err)
	}

	if result.DeletedCount == 0 {
		return common.ErrNotFound
	}

	return nil
}

func (b *BookmarksImpl) Update(bookmark Bookmark) error {
	result, err := b.collection.UpdateByID(context.TODO(), bookmark.ID, bson.D{{"$set", bookmark}})
	if err != nil {
		return fmt.Errorf("failed to update bookmark: %s, error: %w", bookmark.ID, err)
	}

	if result.MatchedCount == 0 {
		return common.ErrNotFound
	}

	return nil
}
