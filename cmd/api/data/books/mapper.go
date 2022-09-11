package books

import (
	"database/sql"

	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
)

func toDBModel(entity *books.Book) *Book {
	return &Book{
		Id:        entity.Id,
		Name:      entity.Name,
		Author:    entity.Author,
		CoverPage: sql.NullString{String: entity.CoverPage, Valid: true},
		Synopsis:  sql.NullString{String: entity.Synopsis, Valid: true},
		Price:     entity.Price,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: sql.NullTime{Time: entity.UpdatedAt, Valid: true},
	}
}

func toDomainModel(entity *Book) *books.Book {
	return &books.Book{
		Id:        entity.Id,
		Name:      entity.Name,
		Author:    entity.Author,
		CoverPage: entity.CoverPage.String,
		Synopsis:  entity.Synopsis.String,
		Price:     entity.Price,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt.Time,
	}
}
