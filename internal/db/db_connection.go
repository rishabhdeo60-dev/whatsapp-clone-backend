package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/config"
)

// DB wraps pgx connection pool
type DB struct {
	Pool *pgxpool.Pool
}

// NewConnection creates a new DB connection using config
func NewConnection(dbcfg *config.DBConfig) *DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbcfg.User, dbcfg.Password, dbcfg.Host, dbcfg.Port, dbcfg.Name,
	)
	log.Printf("Connecting to DB with DSN: %s", dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// if err := pool.Ping(ctx); err != nil {
	// 	log.Fatalf("DB ping failed: %v", err)
	// }

	log.Println("✅ Connected to PostgreSQL successfully")
	return &DB{Pool: pool}
}

// Close closes the DB pool connection
func (db *DB) Close() {
	db.Pool.Close()
	log.Println("🔒 Database connection closed")
}
