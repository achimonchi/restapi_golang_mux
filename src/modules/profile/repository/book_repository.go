package repository

import (
	"database/sql"

	"github.com/achimonchi/belajar_restapi_mux/src/modules/profile/model"
)

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll() (model.Books, error) {
	query := `SELECT * FROM books`

	var books model.Books
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book model.Book

		err = rows.Scan(
			&book.ID, &book.Isbn, &book.Title, &book.Author,
			&book.CreatedAt, &book.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
