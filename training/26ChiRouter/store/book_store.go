package store

import (
	"chirouter/models"
	"sync"
	"time"
)

/*BookStore struct with sync.RWMutex for thread safety
Methods: GetAll(), GetByID(id), Create(book), Update(id, book), Delete(id)
Auto-incrementing ID counter
Pre-seeded with 3 sample books*/

type BookStore struct {
	mu     sync.RWMutex // it is use for managing concurrent access
	books  map[int]models.Book // this is act like a data table
	nextId int                // this act as a primary key generator
}



func NewBookStore() *BookStore{
	store := &BookStore{
		books: make(map[int]models.Book),
		nextId: 1,
	}

	store.Create(models.Book{
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price: 100,
		Year: 2022,
	})

	store.Create(models.Book{
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price: 100,
		Year: 2022,
	})

	store.Create(models.Book{
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price: 100,
		Year: 2022,
	})

	return store
}


//getall 
func (s *BookStore) GetAll() []models.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	books := make([]models.Book, 0, len(s.books))

	for _, book := range s.books {
		books = append(books, book)
	}

	return books
}

//GetByID
func (sb *BookStore) GetByID(id int) models.Book {
	sb.mu.RLock()
	defer sb.mu.RUnlock()

	book , exists := sb.books[id]
	if !exists {
		return models.Book{}
	}
	return book
}

//CREATE

func (sb *BookStore) Create(book models.Book) models.Book {
	sb.mu.Lock() // We use a Write Lock because we are changing data
	defer sb.mu.Unlock()

	book.Id = sb.nextId
	book.CreatedAt = time.Now()
	book.UpdatedAt = book.CreatedAt

	sb.books[book.Id] = book
	sb.nextId++

	return book
}

//Update 

func (sb * BookStore) Update (Id int , book models.Book) models.Book{
	sb.mu.Lock()
	defer sb.mu.Unlock()

	_, exists := sb.books[Id]
	if !exists{
		return models.Book{}
	}
	book.UpdatedAt = time.Now()
	book.Id = Id

	sb.books[Id] = book
	return book
}

//Delete

func (sb *BookStore) Delete(id int) {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	_, exists := sb.books[id]
	if !exists {
		return 
	}
	delete(sb.books, id)
}