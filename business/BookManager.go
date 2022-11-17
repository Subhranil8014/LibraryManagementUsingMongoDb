package business

import (
	"librarySystem/dao"

	"librarySystem/models"

	"librarySystem/interfaces"

	"github.com/sirupsen/logrus"
)

type BookManager struct {
	interfaces.IBookManager
}

// func (manager BookManager) GetAllBooks() []models.Book {
// 	res := dao.GetAllBooks()

// 	return res
// }

func (manager BookManager) AddBook(userid string, bookObj models.Book) (models.ErrorMsg, models.Book) {

	res, err := dao.AddBook(userid, bookObj)

	if err.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "AddBook",
			"UserId":     userid,
			"StatusCode": err.Code,
		}).Error(err.Message)
	}
	return err, res
}

func (manager BookManager) GetAvailableBooks(userid string) (models.ErrorMsg, []models.Book) {

	err, resp := dao.GetAvailableBooks(userid)

	if err.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "GetAvailableBooks",
			"StatusCode": err.Code,
		}).Error(err.Message)
	}
	return err, resp
}

func (manager BookManager) GetBooksByNameOrAuthorOrGenre(userid string, bookname string, authorname string, genre string) (models.ErrorMsg, []models.Book) {
	err, resp := dao.GetBooksByNameOrAuthorOrGenre(userid, bookname, authorname, genre)

	if err.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByNameOrAuthorOrGenre",
			"StatusCode": err.Code,
		}).Error(err.Message)
	}
	return err, resp
}

func (manager BookManager) GetBooksByUserName(userid string, lendername string) (models.ErrorMsg, []models.Book) {
	err, resp := dao.GetBooksByUserName(userid, lendername)

	if err.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByUserName",
			"StatusCode": err.Code,
		}).Error(err.Message)
	}
	return err, resp
}

func (manager BookManager) BorrowBook(userid string, bookname string, username string, bookrating float64) models.ErrorMsg {
	res := dao.BorrowBook(userid, bookname, username, bookrating)

	if res.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "BorrowBook",
			"StatusCode": res.Code,
			"UserId":     userid,
		}).Error(res.Message)

	}
	return res
}

func (manager BookManager) GetBooksByStatus(userid string, status string) (models.ErrorMsg, []models.Book) {
	err, resp := dao.GetBooksByStatus(userid, status)

	if err.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByStatus",
			"StatusCode": err.Code,
		}).Error(err.Message)
	}
	return err, resp
}

func (manager BookManager) ReturnBook(userid string, bookname string, lendername string, bookrating float64) models.ErrorMsg {
	res := dao.ReturnBook(userid, bookname, lendername, bookrating)

	if res.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "ReturnBook",
			"StatusCode": res.Code,
			"UserId":     userid,
		}).Error(res.Message)

	}
	return res
}

func (manager BookManager) RemoveBook(userid string, bookid string) models.ErrorMsg {
	res := dao.RemoveBook(userid, bookid)

	if res.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "RemoveBook",
			"StatusCode": res.Code,
			"UserId":     userid,
		}).Error(res.Message)

	}
	return res
}

func (manager BookManager) RateBook(userid string, bookid string, bookrating float64) models.ErrorMsg {
	res := dao.RateBook(userid, bookid, bookrating)

	if res.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "RateBook",
			"StatusCode": res.Code,
			"UserId":     userid,
		}).Error(res.Message)

	}
	return res
}

// func (manager UserManager) RateBook(userid string, bookid string, bookrating float64) models.ErrorMsg {
// 	fmt.Println("hey...")
// 	res := dao.RateBook(userid, bookid, bookrating)

// 	if res.Message != "" {
// 		logrus.WithFields(logrus.Fields{
// 			"Api":        "RateBook",
// 			"StatusCode": res.Code,
// 			"UserId":     userid,
// 		}).Error(res.Message)

// 	}
// 	return res
// }
