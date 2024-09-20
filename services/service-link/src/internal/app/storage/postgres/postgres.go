package postgres

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

func link() string {
	return fmt.Sprintf("postgres://postgres:%s@link-db/postgres",
		os.Getenv("POSTGRES_PASSWORD"),
	)
}

func NewConnectDB() (*DB, error) {
	poolConfig, err := pgxpool.ParseConfig(link())
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &DB{Mutex: sync.Mutex{}, Pool: pool}, nil
}

func CreateTableOrders(db *DB) error {
	conn, err := db.Pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	func() {
		conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS links ()")
		conn.Exec(context.Background(), "ALTER TABLE links ADD link_id SERIAL PRIMARY KEY NOT NULL")
		conn.Exec(context.Background(), "ALTER TABLE links ADD link_addr_in VARCHAR(50) NOT NULL")
		conn.Exec(context.Background(), "ALTER TABLE links ADD link_addr_out VARCHAR(50) NOT NULL")
	}()
	func() {
		conn.Exec(context.Background(), "CREATE INDEX idx_link_addr_in ON links (link_addr_in)")
	}()

	return nil
}
