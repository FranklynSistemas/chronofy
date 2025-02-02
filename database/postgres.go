package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// it's important to import the postgres drivers so that we can use the postgres database
	"github.com/FranklynSistemas/chronofy/internal/models"
	_ "github.com/lib/pq"
)

// PostgresRepository is a struct that will hold the database connection
type PostgresRepository struct {
	db *sql.DB
}

// Instance of the PostgresRepository
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	fmt.Printf("db: %v", db)
	fmt.Printf("err: %v", err)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (repo *PostgresRepository) GetEvents(ctx context.Context, query models.QueryParams) ([]models.Event, error) {
	var count int
	err := repo.db.QueryRow("SELECT COUNT(*) FROM events").Scan(&count)
	if err != nil {
		log.Println("Count Query Error:", err)
	} else {
		log.Println("Total rows in 'events' table:", count)
	}
	var exists bool
	err = repo.db.QueryRow("SELECT EXISTS (SELECT 1 FROM events)").Scan(&exists)
	log.Println("Does table have data?", exists)
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, external_id, payload, created_at FROM events")
	if err != nil {
		log.Println("Query Error:", err)
		return nil, err
	}
	fmt.Println("rows", rows)
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var events []models.Event
	for rows.Next() {
		var event = models.Event{}
		// Scan the result into the user struct
		err := rows.Scan(&event.ID, &event.Name, &event.ExternalID, &event.Payload, &event.CreatedAt)
		if err == nil {
			events = append(events, event)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

// Pending start using this repository on when using the database provider on the fetch-data-handler.go
