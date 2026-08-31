package models

import (
	"errors"
	"time"
)
type Book struct {
	Id        int    `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required,min=2,max=50"`
	Author    string `json:"author" validate:"required"`
	Price     int    `json:"price" validate:"required,gt=0"` //gt is range validate
	Year      int    `json:"year" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// validation logic
func (b *Book) Validate() error {
	if b.Title == "" {
		return errors.New("title is required")
	}

	if len(b.Title) < 2 {
		return errors.New("title must be at least 2 characters")
	}

	if b.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if b.Year < 1000 {
		return errors.New("invalid publication year")
	}

	return nil
}










/*
React
  ↓
HTTP Request
  ↓
Request DTO validation       ← "Is this input structurally valid?"
  ↓
Service validation           ← "Is this operation allowed?"
  ↓
Domain/model rules            ← "Is this state valid?"
  ↓
Repository
  ↓
PostgreSQL
*/
