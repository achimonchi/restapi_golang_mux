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
	query := `SELECT id, isbn, author, title, createdat, updatedat FROM books`

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

func (r *bookRepository) FindByID(id int) (*model.Book, error) {
	query := `SELECT * FROM books WHERE "id"=$1`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var book model.Book

	defer statement.Close()

	err = statement.QueryRow(id).Scan(
		&book.ID, &book.Isbn, &book.Title, &book.Author,
		&book.CreatedAt, &book.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *bookRepository) Save(book *model.Book) error {
	query := `
		INSERT INTO books ("id","isbn","title","author","createdat","updatedat")
		VALUES($1,$2,$3,$4,$5,$6)
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(book.ID, book.Isbn, book.Title, book.Author, book.CreatedAt, book.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *bookRepository) Update(id int, book *model.Book) error {
	query := `
		UPDATE books
		SET "isbn"=$1, "title"=$2, "author"=$3, "updatedat"=$4
		WHERE "id" = $5
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(book.Isbn, book.Title, book.Author, book.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *bookRepository) Delete(id int) error {
	query := `DELETE FROM books WHERE "id"=$1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
