package postgres

import (
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Mutex sync.Mutex
	Pool  *pgxpool.Pool
}
