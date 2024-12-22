package mpostgres

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type PostgresManager struct {
	db *sql.DB
}

var instance *PostgresManager
var once sync.Once

// Connect
func Connect(connectionString string, maxIdle, maxOpen int) (*PostgresManager, func()) {
	once.Do(func() {
		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			log.Fatalf("Failed to connect to Postgres: %v", err)
		}

		// Configure connection pool
		db.SetMaxIdleConns(maxIdle)
		db.SetMaxOpenConns(maxOpen)
		db.SetConnMaxLifetime(time.Minute * 5)

		if err := db.Ping(); err != nil {
			log.Fatalf("Failed to ping Postgres: %v", err)
		}

		instance = &PostgresManager{db: db}
	})
	return instance, close
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	if instance == nil {
		log.Fatal("Call Connect() first.")
	}
	return instance.db
}

// Close closes the database connection
func close() {
	if instance != nil && instance.db != nil {
		_ = instance.db.Close()
	}
}
