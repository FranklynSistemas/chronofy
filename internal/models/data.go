package models

import "time"

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
