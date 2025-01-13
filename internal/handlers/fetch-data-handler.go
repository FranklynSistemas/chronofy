package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/FranklynSistemas/chronofy/internal/models"
	"github.com/FranklynSistemas/chronofy/internal/providers"
	"github.com/FranklynSistemas/chronofy/internal/services"
	"github.com/gin-gonic/gin"
)

func FetchDataHandler(c *gin.Context) {
	// Parse query parameters.
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	order := c.DefaultQuery("order", "asc") // Default to ascending order.
	providerNames := c.Query("providers")   // Comma-separated provider names.

	// Validate and parse dates.
	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use RFC3339 format."})
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use RFC3339 format."})
		return
	}

	// Construct query parameters.
	params := models.QueryParams{
		StartDate: startDate,
		EndDate:   endDate,
		Order:     order,
	}

	// Filter providers if the `providers` query parameter is provided.
	var selectedProviders []providers.Provider
	allProviders := map[string]providers.Provider{
		"gcp":      &providers.GCPLogsProvider{},
		"database": &providers.DatabaseProvider{},
		"sentry":   &providers.SentryProvider{},
	}

	if providerNames != "" {
		// Split the `providers` query into a list and filter the provider list.
		requestedProviders := strings.Split(providerNames, ",")
		for _, pName := range requestedProviders {
			if provider, exists := allProviders[strings.ToLower(strings.TrimSpace(pName))]; exists {
				selectedProviders = append(selectedProviders, provider)
			}
		}
	} else {
		// No providers specified; use all available providers.
		for _, provider := range allProviders {
			selectedProviders = append(selectedProviders, provider)
		}
	}

	// Fetch data from selected providers.
	data, fetchErr := services.FetchDataFromProviders(params, selectedProviders)
	if fetchErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fetchErr.Error()})
		return
	}

	// Return the normalized and sorted data.
	c.JSON(http.StatusOK, gin.H{"data": data})
}
