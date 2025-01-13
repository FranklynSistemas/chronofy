package providers

import (
	"time"

	"github.com/FranklynSistemas/chronofy/internal/models"
)

type SentryProvider struct{}

type SentryErrorEntry struct { // Example Sentry-specific error structure.
	ErrorID   string
	Message   string
	Timestamp time.Time
	Severity  string
}

func (p *SentryProvider) FetchData(params models.QueryParams) ([]interface{}, error) {
	// Simulate fetching raw data from Sentry.
	rawData := []interface{}{
		SentryErrorEntry{ErrorID: "3", Message: "Sample error", Timestamp: time.Now(), Severity: "CRITICAL"},
	}
	return rawData, nil
}

func (p *SentryProvider) Normalize(rawData []interface{}) []models.Data {
	var normalized []models.Data
	for _, entry := range rawData {
		errorEntry := entry.(SentryErrorEntry) // Type assertion to SentryErrorEntry.
		normalized = append(normalized, models.Data{
			ID:        errorEntry.ErrorID,
			Source:    "Sentry",
			Timestamp: errorEntry.Timestamp,
			Type:      "event",
			Body: map[string]interface{}{
				"message":  errorEntry.Message,
				"severity": errorEntry.Severity,
			},
		})
	}
	return normalized
}
