package interfaces

import (
	"librarySystem/models"
)

type IBookManager interface {
	// GetAllBooks() []models.Book
	AddBook(userid string, bookObj models.Book) (models.ErrorMsg, models.Book)
	GetAvailableBooks(userid string) (models.ErrorMsg, []models.Book)
	GetBooksByNameOrAuthorOrGenre(userid string, bookname string, authorname string, genre string) (models.ErrorMsg, []models.Book)
	GetBooksByUserName(userid string, lendername string) (models.ErrorMsg, []models.Book)
	BorrowBook(userid string, bookname string, username string, bookrating float64) models.ErrorMsg
	GetBooksByStatus(userid string, status string) (models.ErrorMsg, []models.Book)
	ReturnBook(userid string, bookname string, lendername string, bookrating float64) models.ErrorMsg
	RemoveBook(userid string, bookid string) models.ErrorMsg
	// RateBook(userid string, bookid string, bookrating float64) models.ErrorMsg
}
