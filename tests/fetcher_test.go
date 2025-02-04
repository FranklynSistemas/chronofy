package tests

import (
	"context"
	"testing"
	"time"

	"github.com/FranklynSistemas/chronofy/internal/models"
	"github.com/FranklynSistemas/chronofy/internal/providers"
	"github.com/FranklynSistemas/chronofy/internal/services"
)

type MockProvider struct{}

type MockData struct {
	ID        string
	Body      map[string]interface{}
	Timestamp time.Time
}

// TODO: fix the test
// Look if there is possible connect with gcp logs

func TestFetchDataFromProviders(t *testing.T) {
	mockProvider := &MockProvider{}
	providers := []providers.Provider{mockProvider}
	params := models.QueryParams{
		StartDate: time.Now().Add(-24 * time.Hour),
		EndDate:   time.Now(),
	}
	context := context.Background()
	data, err := services.FetchDataFromProviders(context, params, providers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(data) == 0 {
		t.Fatal("expected data, got none")
	}
}

func TestNormalizeData(t *testing.T) {
	rawData := []interface{}{
		MockData{ID: "mock1", Timestamp: time.Now(), Body: map[string]interface{}{"key1": "value1"}},
		MockData{ID: "mock2", Timestamp: time.Now(), Body: map[string]interface{}{"key2": "value2"}},
	}

	mockProvider := &MockProvider{}
	data := mockProvider.Normalize(rawData)
	if len(data) != 2 {
		t.Fatalf("expected 2 elements, got %d", len(data))
	}
	// check if the data is normalized correctly
	for i, d := range data {
		mockData := rawData[i].(MockData)
		if d.ID != mockData.ID {
			t.Errorf("expected ID %s, got %s", mockData.ID, d.ID)
		}
		if d.Timestamp != mockData.Timestamp {
			t.Errorf("expected Timestamp %v, got %v", mockData.Timestamp, d.Timestamp)
		}
		if d.Body["key1"] != mockData.Body["key1"] {
			t.Errorf("expected Body %v, got %v", mockData.Body, d.Body)
		}
	}
}

func (p *MockProvider) FetchData(ctx context.Context, params models.QueryParams) ([]interface{}, error) {
	data := []MockData{
		{ID: "mock1", Timestamp: time.Now(), Body: map[string]interface{}{"key": "value"}},
	}
	interfaceData := make([]interface{}, len(data))
	for i, v := range data {
		interfaceData[i] = v
	}
	return interfaceData, nil
}

func (p *MockProvider) Normalize(rawData []interface{}) []models.Data {
	var normalized []models.Data
	for _, entry := range rawData {
		data := entry.(MockData) // Type assertion to DBLogEntry.
		normalized = append(normalized, models.Data{
			ID:        data.ID,
			Source:    "Database",
			Timestamp: data.Timestamp,
			Type:      "mock",
			Body:      data.Body,
		})
	}
	return normalized
}
