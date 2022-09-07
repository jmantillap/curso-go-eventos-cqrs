package repository

import (
	"context"

	"platzi.com/go/cqrs/models"
)

type Repository interface {
	Close()
	InsertFeed(ctx context.Context, feed *models.Feed) error
	ListFeeds(ctx context.Context) ([]*models.Feed, error)
}

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

func Close() {
	repository.Close()
}
