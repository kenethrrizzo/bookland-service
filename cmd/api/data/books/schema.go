package books

import (
	"database/sql"
	"time"
)

type Book struct {
	Id        int            `db:"Id"`
	Name      string         `db:"Name"`
	Author    int            `db:"Author"`
	CoverPage sql.NullString `db:"CoverPage"`
	Synopsis  sql.NullString `db:"Synopsis"`
	Price     float64        `db:"Price"`
	CreatedAt time.Time      `db:"CreatedAt"`
	UpdatedAt sql.NullTime   `db:"UpdatedAt"`
}
