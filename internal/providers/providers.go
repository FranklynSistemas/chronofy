package providers

import (
	"github.com/FranklynSistemas/chronofy/internal/models"
)

type Provider interface {
	FetchData(params models.QueryParams) ([]interface{}, error) // Returns raw provider-specific data.
	Normalize(rawData []interface{}) []models.Data              // Normalizes raw data into Data.
}
