package interfaces

import (
	"net/http"
)

type IBookService interface {
	// GetAllBooks(response http.ResponseWriter, request *http.Request)
	AddBook(response http.ResponseWriter, request *http.Request)
	GetAvailableBooks(response http.ResponseWriter, request *http.Request)
	GetBooksByNameOrAuthorOrGenre(response http.ResponseWriter, request *http.Request)
	GetBooksByUserName(response http.ResponseWriter, request *http.Request)
	BorrowBook(response http.ResponseWriter, request *http.Request)
	GetBooksByStatus(response http.ResponseWriter, request *http.Request)
	ReturnBook(response http.ResponseWriter, request *http.Request)
	RemoveBook(response http.ResponseWriter, request *http.Request)
	RateBook(response http.ResponseWriter, request *http.Request)
}
