package pg

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PG struct {
	DB *pgxpool.Pool
}

func New(ctx context.Context, dbUrl string) *PG {
	pg, err := connect(dbUrl)
	if err != nil {
		slog.ErrorContext(context.Background(), "error connect to db", "error", err)
		log.Panicf("Failed to connect to database: %v", err)
	}

	log.Println("ping database...")
	if err = pg.Ping(ctx); err != nil {
		log.Panicf("Failed to ping database: %v", err)
		os.Exit(1)
	}
	log.Println("ping db succeed!")

	return &PG{
		DB: pg,
	}
}

func connect(dbUrl string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return nil, err
	}

	// Don't close the pool here - it should remain open for the application
	// The original defer dbpool.Close() was causing the "closed pool" error

	return dbpool, nil
}
