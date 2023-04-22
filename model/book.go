package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID          int64     `json:"id" gorm:"column:id"`
	Title       string    `json:"title" gorm:"column:title" validate:"required,min=3,max=100"`
	Author      string    `json:"author" gorm:"column:author" validate:"required,min=3,max=100"`
	Description string    `json:"description" gorm:"column:description" validate:"required,min=3,max=1000"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type BookRequest struct {
	Title       string `json:"title" example:"Test Book"`
	Author      string `json:"author" example:"Test Author"`
	Description string `json:"description" example:"Test Desc"`
}

func (m *Book) TableName() string {
	return "public.books"
}

func (e Book) Validation() error { // custom validation
	return validation.ValidateStruct(&e,
		validation.Field(&e.Title, validation.Required, validation.Length(3, 100)),
		validation.Field(&e.Author, validation.Required, validation.Length(3, 100)),
		validation.Field(&e.Description, validation.Required, validation.Length(3, 1000)))
}
