package services

import (
	"sort"
	"sync"

	"github.com/FranklynSistemas/chronofy/internal/models"
	"github.com/FranklynSistemas/chronofy/internal/providers"
)

func FetchDataFromProviders(params models.QueryParams, providerList []providers.Provider) ([]models.Data, error) {
	var wg sync.WaitGroup
	results := make(chan []models.Data, len(providerList)) // Channel to collect results.
	errors := make(chan error, len(providerList))          // Channel to collect errors.

	// Launch a goroutine for each provider.
	for _, provider := range providerList {
		wg.Add(1)
		go func(p providers.Provider) {
			defer wg.Done()

			// Fetch raw data.
			rawData, err := p.FetchData(params)
			if err != nil {
				errors <- err
				return
			}

			// Normalize data.
			normalizedData := p.Normalize(rawData)
			results <- normalizedData
		}(provider)
	}

	// Close channels once all goroutines are complete.
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	// Collect results and errors.
	var allData []models.Data
	for data := range results {
		allData = append(allData, data...)
	}

	// Check for errors.
	if len(errors) > 0 {
		return allData, <-errors // Return the first error encountered.
	}

	// Sort the results based on the order parameter.
	sortData(&allData, params.Order)

	return allData, nil
}

// sortData sorts the data slice in ascending or descending order by Timestamp.
func sortData(data *[]models.Data, order string) {
	sort.SliceStable(*data, func(i, j int) bool {
		if order == "desc" {
			return (*data)[i].Timestamp.After((*data)[j].Timestamp)
		}
		return (*data)[i].Timestamp.Before((*data)[j].Timestamp)
	})
}
