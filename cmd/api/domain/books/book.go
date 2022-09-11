package books

import "time"

type Book struct {
	Id          int
	Name        string
	Author      int
	CoverPage   string
	Synopsis    string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
