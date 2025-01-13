package providers

import (
	"time"

	"github.com/FranklynSistemas/chronofy/internal/models"
)

type DatabaseProvider struct{}

type DBLogEntry struct { // Example database-specific log structure.
	RecordID   string
	Query      string
	ExecutedAt time.Time
	Status     string
}

func (p *DatabaseProvider) FetchData(params models.QueryParams) ([]interface{}, error) {
	// Simulate fetching raw data from a database.
	rawData := []interface{}{
		DBLogEntry{RecordID: "2", Query: "SELECT * FROM logs", ExecutedAt: time.Now(), Status: "SUCCESS"},
	}
	return rawData, nil
}

func (p *DatabaseProvider) Normalize(rawData []interface{}) []models.Data {
	var normalized []models.Data
	for _, entry := range rawData {
		record := entry.(DBLogEntry) // Type assertion to DBLogEntry.
		normalized = append(normalized, models.Data{
			ID:        record.RecordID,
			Source:    "Database",
			Timestamp: record.ExecutedAt,
			Type:      "db",
			Body: map[string]interface{}{
				"query":  record.Query,
				"status": record.Status,
			},
		})
	}
	return normalized
}
