package books

import "time"

type Book struct {
	ID        int
	Name      string
	Author    string
	CoverPage string
	Synopsis  string
	Price     float64
	Genres    []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
