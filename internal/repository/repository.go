package repository

import (
	"context"

	"github.com/FranklynSistemas/chronofy/internal/models"
)

type Repository interface {
	GetEvents(ctx context.Context, query models.QueryParams) ([]models.Event, error)
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo
}

func GetEvents(ctx context.Context, query models.QueryParams) ([]models.Event, error) {
	return implementation.GetEvents(ctx, query)
}
