package providers

import (
	"context"

	"github.com/FranklynSistemas/chronofy/internal/models"
)

type Provider interface {
	FetchData(ctx context.Context, params models.QueryParams) ([]interface{}, error) // Returns raw provider-specific data.
	Normalize(rawData []interface{}) []models.Data                                   // Normalizes raw data into Data.
}
