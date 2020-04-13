package model

import "time"

type Book struct {
	ID        int    `json:"id"`
	Isbn      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Books []Book

func NewBook() *Book {
	return &Book{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
