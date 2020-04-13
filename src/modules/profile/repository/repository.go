package repository

import "github.com/achimonchi/belajar_restapi_mux/src/modules/profile/model"

type BookRepository interface {
	Save(*model.Book) error
	Update(int, *model.Book) error
	Delete(int) error
	FindAll() (model.Books, error)
	FindByID(int) (*model.Book, error)
}
