package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"chirouter/models"
	"chirouter/store"

	"github.com/go-chi/chi"
)

// BookHandler holds the store dependency
type BookHandler struct {
	Store *store.BookStore
}

func NewBookHandler(s *store.BookStore) *BookHandler {
	return &BookHandler{Store: s}
}

// GetAllBooks godoc
// @Summary Get all books
// @Description get all books in the store
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := h.Store.GetAll()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBookByID godoc
// @Summary Get a book by ID
// @Description get book by ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid ID format"
// @Failure 404 {string} string "Book not found"
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	book := h.Store.GetByID(id)
	if book.Id == 0 { // If id is 0, it means it wasn't found in our store
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook godoc
// @Summary Create a new book
// @Description add a new book to the store
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book object"
// @Success 201 {object} models.Book
// @Failure 400 {string} string "Invalid JSON payload"
// @Router /books [post]
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	
	// Decode JSON from request body
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Run our struct validations
	if err := book.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBook := h.Store.Create(book)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description update book by ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book object"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid JSON payload"
// @Failure 404 {string} string "Book not found"
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := book.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook := h.Store.Update(id, book)
	if updatedBook.Id == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description delete book by ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid ID format"
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	h.Store.Delete(id)
	
	w.WriteHeader(http.StatusNoContent) // 204 means deleted successfully with no body
}