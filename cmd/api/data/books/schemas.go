package books

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        int            `db:"Id"`
	Name      string         `db:"Name"`
	Author    int            `db:"Author"`
	Genres    string         `db:"Genres"`
	CoverPage sql.NullString `db:"CoverPage"`
	Synopsis  sql.NullString `db:"Synopsis"`
	Price     float64        `db:"Price"`
	CreatedAt time.Time      `db:"CreatedAt"`
	UpdatedAt sql.NullTime   `db:"UpdatedAt"`
}

type BookGenre struct {
	BookID    int    `db:"BookId"`
	GenreCode string `db:"GenreCode"`
}
