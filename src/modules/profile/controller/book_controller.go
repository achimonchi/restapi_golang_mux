package controller

import (
	"fmt"
	"strconv"

	"github.com/achimonchi/belajar_restapi_mux/src/modules/profile/model"
	"github.com/achimonchi/belajar_restapi_mux/src/modules/profile/repository"
)

func GetAll(repo repository.BookRepository) (model.Books, error) {
	books, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return books, nil
}

func GetByID(id string, repo repository.BookRepository) (*model.Book, error) {
	idInt, _ := strconv.Atoi(id)
	book, err := repo.FindByID(idInt)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func SaveBook(p *model.Book, repo repository.BookRepository) error {
	err := repo.Save(p)
	if err != nil {
		return err
	}

	fmt.Println("Save Success !")

	return nil
}

func Update(b *model.Book, repo repository.BookRepository) error {
	err := repo.Update(b.ID, b)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Update success !")
	return nil
}

func Delete(id string, repo repository.BookRepository) error {
	ID, _ := strconv.Atoi(id)
	err := repo.Delete(ID)
	if err != nil {
		return err
	}

	fmt.Print("Delete Success !")
	return nil
}
