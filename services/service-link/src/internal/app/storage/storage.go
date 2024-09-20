package storage

import "github.com/aygumov-g/shortlinks-go/services/service-link/src/internal/app/storage/postgres"

func NewStorage() (*postgres.DB, error) {
	db, err := postgres.NewConnectDB()
	if err != nil {
		return nil, err
	}
	if err := postgres.CreateTableOrders(db); err != nil {
		return nil, err
	}

	return db, nil
}
