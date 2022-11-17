package services

import (
	"encoding/json"
	"librarySystem/business"
	"librarySystem/models"
	"strconv"

	"librarySystem/interfaces"

	"librarySystem/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type BookService struct {
	interfaces.IBookService
}

// func (service BookService) GetAllBooks(response http.ResponseWriter, request *http.Request) {

// 	books := business.BookManager{}.GetAllBooks()
// 	if len(books) == 0 {
// 		message := utils.GET_ALL_BOOKS_FAILED
// 		response.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(response).Encode(message)
// 		return
// 	}

// 	response.WriteHeader(http.StatusAccepted)
// 	json.NewEncoder(response).Encode(books)
// }

func (service BookService) AddBook(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	userid := params["userid"]

	logrus.WithFields(logrus.Fields{
		"Api":        "AddBook",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.ADD_BOOK_REQUEST)

	var errorBody models.ErrorBody
	var requestBody models.Book
	json.NewDecoder(request.Body).Decode(&requestBody)

	var resp models.ResponseBookObject

	url, res := business.BookManager{}.AddBook(userid, requestBody)
	if url.Message != "" {
		errorBody.Error = url
		logrus.WithFields(logrus.Fields{
			"Api":        "AddBook",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.ADD_BOOK_FAILED)
		response.WriteHeader(url.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "AddBook",
		"Path":       request.URL.Path,
		"StatusCode": 201,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.ADD_BOOK_SUCCESS)

	resp.Data = res
	resp.Message = "Book added successfully"
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(resp)

	return

}

func (service BookService) GetAvailableBooks(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	userid := params["userid"]

	logrus.WithFields(logrus.Fields{
		"Api":        "GetAvailableBooks",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.GET_AVAILABLE_BOOKS_REQUEST)

	var resp1 []models.Book
	var msg models.ResponseDataString
	var resp2 models.ResponseDataSuccess

	var errorBody models.ErrorBody

	url, res := business.BookManager{}.GetAvailableBooks(userid)

	if url.Message != "" {
		errorBody.Error = url
		logrus.WithFields(logrus.Fields{
			"Api":        "GetAvailableBooks",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.GET_AVAILABLE_BOOKS_FAILED)
		response.WriteHeader(url.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	if len(res) == 0 {
		msg.SuccessMsg = utils.NO_AVAILABLE_BOOKS_TO_SHOW
		resp2.Data = msg
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(resp2)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "GetAvailableBooks",
		"Path":       request.URL.Path,
		"StatusCode": 200,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.GET_AVAILABLE_BOOKS_SUCCESS)
	resp1 = res
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp1)
	return
}

func (service BookService) GetBooksByNameOrAuthorOrGenre(response http.ResponseWriter, request *http.Request) {
	logrus.WithFields(logrus.Fields{
		"Api":        "GetBooksByNameOrAuthorOrGenre",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.GET_BOOKS_BY_NAME_OR_AUTHOR_OR_GENRE_REQUEST)

	params := mux.Vars(request)
	userid := params["userid"]
	bookname := request.URL.Query().Get("bookname")
	authorname := request.URL.Query().Get("authorname")
	genre := request.URL.Query().Get("genre")

	var resp1 []models.Book
	var msg models.ResponseDataString
	var resp2 models.ResponseDataSuccess

	var errorBody models.ErrorBody

	url, res := business.BookManager{}.GetBooksByNameOrAuthorOrGenre(userid, bookname, authorname, genre)

	if url.Message != "" {
		errorBody.Error = url
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByNameOrAuthorOrGenre",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.GET_BOOKS_BY_NAME_OR_AUTHOR_OR_GENRE_FAILED)
		response.WriteHeader(url.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	if len(res) == 0 {
		msg.SuccessMsg = utils.NO_BOOKS_TO_SHOW
		resp2.Data = msg
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(resp2)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "GetBooksByNameOrAuthorOrGenre",
		"Path":       request.URL.Path,
		"StatusCode": 200,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.GET_BOOKS_BY_NAME_OR_AUTHOR_OR_GENRE_SUCCESS)
	resp1 = res
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp1)
	return
}

func (service BookService) GetBooksByUserName(response http.ResponseWriter, request *http.Request) {
	logrus.WithFields(logrus.Fields{
		"Api":        "GetBooksByUserName",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.GET_BOOKS_BY_USERNAME_REQUEST)

	params := mux.Vars(request)
	userid := params["userid"]
	lendername := request.URL.Query().Get("lendername")

	var resp1 []models.Book
	var msg models.ResponseDataString
	var resp2 models.ResponseDataSuccess

	var errorBody models.ErrorBody

	url, res := business.BookManager{}.GetBooksByUserName(userid, lendername)

	if url.Message != "" {
		errorBody.Error = url
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByUserName",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.GET_BOOKS_BY_USERNAME_FAILED)
		response.WriteHeader(url.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	if len(res) == 0 {
		msg.SuccessMsg = utils.NO_BOOKS_TO_SHOW
		resp2.Data = msg
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(resp2)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "GetBooksByUserName",
		"Path":       request.URL.Path,
		"StatusCode": 200,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.GET_BOOKS_BY_USERNAME_SUCCESS)
	resp1 = res
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp1)
	return
}

func (service BookService) BorrowBook(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	userid := params["userid"]
	bookname := params["bookname"]
	username := params["username"]
	bookrating := params["bookrating"]
	rating, _ := strconv.ParseFloat(bookrating, 32)

	logrus.WithFields(logrus.Fields{
		"Api":        "BorrowBook",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.BORROW_BOOK_REQUEST)

	var msg models.ResponseDataString
	var resp models.ResponseDataSuccess
	var errorBody models.ErrorBody

	res := business.BookManager{}.BorrowBook(userid, bookname, username, rating)
	if res.Message != "" {
		errorBody.Error = res
		logrus.WithFields(logrus.Fields{
			"Api":        "BorrowBook",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.BORROW_BOOK_FAILED)
		response.WriteHeader(res.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "BorrowBook",
		"Path":       request.URL.Path,
		"ErrorCode":  200,
		"UserId":     userid,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.BORROW_BOOK_SUCCESS)
	msg.SuccessMsg = utils.SUCCESS
	resp.Data = msg
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp)
}

func (service BookService) GetBooksByStatus(response http.ResponseWriter, request *http.Request) {
	logrus.WithFields(logrus.Fields{
		"Api":        "GetBooksByStatus",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.GET_BOOKS_BY_STATUS_REQUEST)

	params := mux.Vars(request)
	userid := params["userid"]
	status := request.URL.Query().Get("status")

	var resp1 []models.Book
	var msg models.ResponseDataString
	var resp2 models.ResponseDataSuccess

	var errorBody models.ErrorBody

	url, res := business.BookManager{}.GetBooksByStatus(userid, status)

	if url.Message != "" {
		errorBody.Error = url
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByStatus",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.GET_BOOKS_BY_STATUS_FAILED)
		response.WriteHeader(url.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	if len(res) == 0 {
		msg.SuccessMsg = utils.NO_BOOKS_TO_SHOW
		resp2.Data = msg
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(resp2)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "GetBooksByStatus",
		"Path":       request.URL.Path,
		"StatusCode": 200,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.GET_BOOKS_BY_STATUS_SUCCESS)
	resp1 = res
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp1)
	return
}

func (service BookService) ReturnBook(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	userid := params["userid"]
	bookname := params["bookname"]
	lendername := params["lendername"]
	bookrating := params["bookrating"]
	rating, _ := strconv.ParseFloat(bookrating, 32)

	logrus.WithFields(logrus.Fields{
		"Api":        "ReturnBook",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.RETURN_BOOK_REQUEST)

	var msg models.ResponseDataString
	var resp models.ResponseDataSuccess
	var errorBody models.ErrorBody

	res := business.BookManager{}.ReturnBook(userid, bookname, lendername, rating)
	if res.Message != "" {
		errorBody.Error = res
		logrus.WithFields(logrus.Fields{
			"Api":        "ReturnBook",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.RETURN_BOOK_FAILED)
		response.WriteHeader(res.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "ReturnBook",
		"Path":       request.URL.Path,
		"ErrorCode":  200,
		"UserId":     userid,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.RETURN_BOOK_SUCCESS)
	msg.SuccessMsg = utils.RETURN_SUCCESS
	resp.Data = msg
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp)
}

func (service BookService) RemoveBook(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	userid := params["userid"]
	bookid := params["bookid"]

	logrus.WithFields(logrus.Fields{
		"Api":        "RemoveBook",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.REMOVE_BOOK_REQUEST)

	var msg models.ResponseDataString
	var resp models.ResponseDataSuccess
	var errorBody models.ErrorBody

	res := business.BookManager{}.RemoveBook(userid, bookid)
	if res.Message != "" {
		errorBody.Error = res
		logrus.WithFields(logrus.Fields{
			"Api":        "RemoveBook",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.REMOVE_BOOK_FAILED)
		response.WriteHeader(res.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "RemoveBook",
		"Path":       request.URL.Path,
		"ErrorCode":  204,
		"UserId":     userid,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.REMOVE_BOOK_SUCCESS)
	msg.SuccessMsg = utils.REMOVE_SUCCESS
	resp.Data = msg
	response.WriteHeader(http.StatusNoContent)
	json.NewEncoder(response).Encode(resp)
	return
}

func (service BookService) RateBook(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userid := params["userid"]
	bookid := params["bookid"]
	bookrating := params["bookrating"]
	rating, _ := strconv.ParseFloat(bookrating, 32)

	logrus.WithFields(logrus.Fields{
		"Api":        "RateBook",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.RATE_BOOK_REQUEST)

	var msg models.ResponseDataString
	var resp models.ResponseDataSuccess
	var errorBody models.ErrorBody

	res := business.BookManager{}.RateBook(userid, bookid, rating)
	if res.Message != "" {
		errorBody.Error = res
		logrus.WithFields(logrus.Fields{
			"Api":        "RateBook",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.RATE_BOOK_FAILED)
		response.WriteHeader(res.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "RateBook",
		"Path":       request.URL.Path,
		"ErrorCode":  200,
		"UserId":     userid,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.RATE_BOOK_SUCCESS)
	msg.SuccessMsg = utils.SUCCESS_RATE_BOOK
	resp.Data = msg
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp)
	return
}
