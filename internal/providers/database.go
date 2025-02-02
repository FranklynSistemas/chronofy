package providers

import (
	"context"
	"fmt"
	"time"

	"github.com/FranklynSistemas/chronofy/internal/models"
	"github.com/FranklynSistemas/chronofy/internal/repository"
)

type DatabaseProvider struct{}

type DBLogEntry struct { // Example database-specific log structure.
	RecordID   string
	Query      string
	ExecutedAt time.Time
	Status     string
}

func (p *DatabaseProvider) FetchData(ctx context.Context, params models.QueryParams) ([]interface{}, error) {
	// Simulate fetching raw data from a database.
	// rawData := []interface{}{
	// 	DBLogEntry{RecordID: "2", Query: "SELECT * FROM logs", ExecutedAt: time.Now(), Status: "SUCCESS"},
	// }
	fmt.Println("Fetching data from database...")
	rawData, _ := repository.GetEvents(ctx, params) // Call the repository function.
	var result []interface{}
	for _, event := range rawData {
		result = append(result, event)
	}
	fmt.Println("data", result)
	return result, nil
}

func (p *DatabaseProvider) Normalize(rawData []interface{}) []models.Data {
	var normalized []models.Data
	for _, entry := range rawData {
		record := entry.(models.Event) // Type assertion to DBLogEntry.
		normalized = append(normalized, models.Data{
			ID:        fmt.Sprintf("%d", record.ID),
			Source:    "Database",
			Timestamp: record.CreatedAt,
			Type:      "db",
			Body: map[string]interface{}{
				"payload":    record.Payload,
				"externalId": record.ExternalID,
			},
		})
	}
	return normalized
}
