package providers

import (
	"context"
	"time"

	"github.com/FranklynSistemas/chronofy/internal/models"
)

type GCPLogsProvider struct{}

type GCPLogEntry struct { // Example GCP-specific log structure.
	LogID     string
	Message   string
	Timestamp time.Time
	LogLevel  string
}

func (p *GCPLogsProvider) FetchData(ctx context.Context, params models.QueryParams) ([]interface{}, error) {
	// Simulate fetching raw data from GCP logs.
	rawData := []interface{}{
		GCPLogEntry{LogID: "1", Message: "GCP log entry", Timestamp: time.Now(), LogLevel: "INFO"},
	}
	return rawData, nil
}

func (p *GCPLogsProvider) Normalize(rawData []interface{}) []models.Data {
	var normalized []models.Data
	for _, entry := range rawData {
		log := entry.(GCPLogEntry) // Type assertion to GCPLogEntry.
		normalized = append(normalized, models.Data{
			ID:        log.LogID,
			Source:    "GCP",
			Timestamp: log.Timestamp,
			Type:      "log",
			Body: map[string]interface{}{
				"message": log.Message,
				"level":   log.LogLevel,
			},
		})
	}
	return normalized
}
