package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func Connect(dataBaseUrl string) (*sql.DB, error) {
	//watch database pool video of codersgyan
	db, err := sql.Open("pgx", dataBaseUrl)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db,Ping: %w", err)
	}
	return db, nil
}
