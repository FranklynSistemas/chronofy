package tests

import (
	"testing"
	"time"

	"github.com/FranklynSistemas/chronofy/internal/models"
	"github.com/FranklynSistemas/chronofy/internal/services"
)

func TestFetchDataFromProviders(t *testing.T) {
	mockProvider := &MockProvider{}
	providers := []services.Provider{mockProvider}
	params := services.QueryParams{
		StartDate: time.Now().Add(-24 * time.Hour),
		EndDate:   time.Now(),
	}

	data, err := services.FetchDataFromProviders(params, providers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(data) == 0 {
		t.Fatal("expected data, got none")
	}
}

type MockProvider struct{}

func (p *MockProvider) FetchData(params services.QueryParams) ([]models.Data, error) {
	return []models.Data{
		{ID: "mock1", Source: "mock", Timestamp: time.Now(), Type: "mock", Body: map[string]interface{}{"key": "value"}},
	}, nil
}
