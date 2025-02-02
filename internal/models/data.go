package models

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID         int32           `json:"id"`
	Name       string          `json:"name"`
	Payload    json.RawMessage `json:"payload"`
	ExternalID string          `json:"external_id"`
	CreatedAt  time.Time       `json:"created_at"`
}

type Data struct {
	ID        string
	Source    string
	Timestamp time.Time
	Type      string
	Body      map[string]interface{}
}

type QueryParams struct {
	StartDate time.Time
	EndDate   time.Time
	Source    string
	Filters   map[string]string
	Order     string
	Providers []string
}
