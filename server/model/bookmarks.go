package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mymmrac/mymm.gq/server/common"
	"github.com/mymmrac/mymm.gq/server/logger"
)

type ObjectID = primitive.ObjectID

var NilObjectID = primitive.NilObjectID

type BookmarkAddRequest struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Category string `json:"category"`
}

type Bookmark struct {
	ID       ObjectID `json:"id"                 bson:"_id,omitempty"`
	Name     string   `json:"name"               bson:"name"`
	Link     string   `json:"link"               bson:"link"`
	Category string   `json:"category"           bson:"category"`
	IconLink string   `json:"iconLink,omitempty" bson:"iconLink,omitempty"`
}

type Bookmarks interface {
	Bookmarks() ([]Bookmark, error)
	Add(request BookmarkAddRequest) (*Bookmark, error)
	Update(bookmark Bookmark) error
	Delete(id ObjectID) error
}

type BookmarksImpl struct {
	log        logger.Logger
	httpClient *http.Client
	collection *mongo.Collection
}

func NewBookmarks(log logger.Logger, client *mongo.Client) *BookmarksImpl {
	collection := client.Database("mymm").Collection("bookmarks")
	return &BookmarksImpl{
		log:        log,
		httpClient: http.DefaultClient,
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

func (b *BookmarksImpl) Add(request BookmarkAddRequest) (*Bookmark, error) {
	bookmark := Bookmark{
		Name:     request.Name,
		Link:     request.Link,
		Category: request.Category,
	}

	iconLink, err := b.grabIcon(request.Link)
	if err != nil {
		b.log.Errorf("Failed to get icon for %s, error: %s", request.Link, err)
	} else {
		bookmark.IconLink = iconLink
	}

	result, err := b.collection.InsertOne(context.TODO(), bookmark)
	if err != nil {
		return nil, fmt.Errorf("failed to add bookmark: %w", err)
	}

	id, ok := result.InsertedID.(ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to cast bookmark ID: %w", err)
	}

	bookmark.ID = id
	return &bookmark, nil
}

func (b *BookmarksImpl) Delete(id ObjectID) error {
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

const faviconAPI = "https://favicongrabber.com/api/grab/%s"

type iconGrabberResp struct {
	Icons []struct {
		Src string `json:"src"`
	} `json:"icons"`
}

func (b *BookmarksImpl) grabIcon(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	domain := u.Hostname()

	resp, err := b.httpClient.Get(fmt.Sprintf(faviconAPI, domain))
	if err != nil {
		return "", fmt.Errorf("failed to call icon grabber: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed grab icon, status code: %d", resp.StatusCode)
	}

	var icons iconGrabberResp
	if err = json.NewDecoder(resp.Body).Decode(&icons); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(icons.Icons) == 0 {
		return "", fmt.Errorf("failed to grab icons: %w", common.ErrNotFound)
	}

	return icons.Icons[0].Src, nil
}
