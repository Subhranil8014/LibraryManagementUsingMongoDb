package main

import (
	"fmt"

	"librarySystem/services"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	fmt.Println("Initialize Router")

	router := mux.NewRouter()

	// router.HandleFunc("/library/books", services.BookService{}.GetAllBooks).Methods("GET")
	router.HandleFunc("/library/user", services.UserService{}.CreateUser).Methods("POST")
	router.HandleFunc("/library/{userid}/addBook", services.BookService{}.AddBook).Methods("POST")
	router.HandleFunc("/library/availableBooks/{userid}", services.BookService{}.GetAvailableBooks).Methods("GET")
	router.HandleFunc("/library/{userid}", services.BookService{}.GetBooksByNameOrAuthorOrGenre).Methods("GET")
	router.HandleFunc("/library/username/{userid}", services.BookService{}.GetBooksByUserName).Methods("GET")
	router.HandleFunc("/library/borrow/{userid}/{bookname}/{username}/{bookrating}", services.BookService{}.BorrowBook).Methods("GET")
	router.HandleFunc("/library/books/{userid}", services.BookService{}.GetBooksByStatus).Methods("GET")
	router.HandleFunc("/library/return/{userid}/{bookname}/{lendername}/{bookrating}", services.BookService{}.ReturnBook).Methods("GET")
	router.HandleFunc("/library/remove/{userid}/{bookid}", services.BookService{}.RemoveBook).Methods("DELETE")
	router.HandleFunc("/library/rateuser/{userid}/{username}/{userrating}", services.UserService{}.RateUser).Methods("GET")
	router.HandleFunc("/library/ratebook/{userid}/{bookid}/{bookrating}", services.BookService{}.RateBook).Methods("GET")
	return router
}
