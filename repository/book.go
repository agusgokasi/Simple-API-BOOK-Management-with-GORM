package repository

import (
	"ninth-learn/model"
)

// clean architectures -> handler->service->repo

// interface Book
type BookRepo interface {
	GetBooks() ([]model.Book, error)
	CreateBook(in model.Book) (res model.Book, err error)
	GetBookById(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) GetBooks() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	result := r.db.Create(&in)
	if result.Error != nil {
		return res, result.Error
	}

	return in, nil
}

func (r Repo) GetBookById(id int64) (res model.Book, err error) {
	if err := r.db.Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) UpdateBook(in model.Book) (res model.Book, err error) {
	// Update book record in the database
	err = r.db.Model(&model.Book{}).Where("id = ?", in.ID).Updates(model.Book{
		Title:       in.Title,
		Author:      in.Author,
		Description: in.Description,
	}).Error
	if err != nil {
		return res, err
	}

	res = in
	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	// Find the book to delete
	book := model.Book{}
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}

	// Delete the book
	if err := r.db.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
