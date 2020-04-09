package repository

import "github.com/achimonchi/belajar_restapi_mux/src/modules/profile/model"

type BookRepository interface {
	Save(*model.Book) error
	Update(string, *model.Book) error
	delete(string) error
	FindAll() (*model.Books, error)
	FindByID(string) (*model.Book, error)
}
