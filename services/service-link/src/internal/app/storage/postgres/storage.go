package postgres

import (
	"context"
	"fmt"

	"github.com/aygumov-g/shortlinks-go/services/service-link/src/internal/app/web/home"
)

func (db *DB) LinkCreate(link home.Link) (home.Link, error) {
	conn, err := db.Pool.Acquire(context.Background())
	if err != nil {
		return home.Link{}, fmt.Errorf("%s", "Server storage connect error")
	}
	defer conn.Release()
	row := conn.QueryRow(
		context.Background(), "INSERT INTO links (link_addr_in, link_addr_out) VALUES ($1, $2) RETURNING link_id, link_addr_in, link_addr_out",
		link.LinkAddrIn,
		link.LinkAddrOut,
	)
	var linkOut home.Link
	if err := row.Scan(&linkOut.LinkId, &linkOut.LinkAddrIn, &linkOut.LinkAddrOut); err != nil {
		return home.Link{}, fmt.Errorf("%s", "Failed to create object link")
	}

	return linkOut, nil
}

func (db *DB) LinkSearch(link_addr_in string) (home.Link, error) {
	conn, err := db.Pool.Acquire(context.Background())
	if err != nil {
		return home.Link{}, fmt.Errorf("%s", "Server storage connect error")
	}
	defer conn.Release()
	row := conn.QueryRow(
		context.Background(), "SELECT link_id, link_addr_in, link_addr_out FROM links WHERE link_addr_in = $1",
		link_addr_in,
	)
	var linkOut home.Link
	if err := row.Scan(&linkOut.LinkId, &linkOut.LinkAddrIn, &linkOut.LinkAddrOut); err != nil {
		return home.Link{}, fmt.Errorf("%s", "Failed to search link")
	}

	return linkOut, nil
}
