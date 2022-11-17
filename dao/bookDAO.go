package dao

import (
	"context"
	"fmt"
	database "librarySystem/dbUtils"
	"librarySystem/models"
	"librarySystem/utils"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func GetAllBooks() []models.Book {

// 	var books []models.Book

// 	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	cursor, err := bookCollection.Find(ctx, bson.M{})

// 	if err != nil {
// 		logrus.WithFields(logrus.Fields{
// 			"Api":        "MongoGetAllBooks",
// 			"StatusCode": 500,
// 		}).Info(err.Error())
// 	}

// 	for cursor.Next(ctx) {
// 		var book models.Book
// 		fmt.Println(book)
// 		err := cursor.Decode(&book)
// 		if err != nil {
// 			logrus.WithFields(logrus.Fields{
// 				"Api":        "MongoGetAllBooks",
// 				"StatusCode": 500,
// 			}).Error(err.Error())
// 		}

// 		books = append(books, book)

// 	}

// 	defer cursor.Close(ctx)

// 	return books

// }

func AddBook(userid string, bookObj models.Book) (models.Book, models.ErrorMsg) {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userObj models.User
	var errorMsg models.ErrorMsg

	err := userCollection.FindOne(ctx, bson.M{"userid": userid}).Decode(&userObj)

	if err != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.USERID_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "AddBook",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err.Error())
		return bookObj, errorMsg
	}

	bookObj.BookId = generateUUID()
	loc, _ := time.LoadLocation("Asia/Kolkata")
	currentTime := time.Now().In(loc)
	t := currentTime.Format(time.RFC850)
	bookObj.CreatedAt = t
	bookObj.BookRating = 0
	bookObj.Status = "ADDED"
	bookObj.UserId = userid
	bookObj.LenderName = userObj.UserName
	bookObj.LenderRating = userObj.UserRating

	userObj.Books = append(userObj.Books, bookObj)
	userCollection.FindOneAndReplace(ctx, bson.M{"userid": userid}, userObj)

	_, err1 := bookCollection.InsertOne(ctx, bookObj)

	if err1 != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.ADD_BOOK_FAILED
		logrus.WithFields(logrus.Fields{
			"Api":        "AddBook",
			"StatusCode": 500,
			"ID":         userid,
		}).Error(err.Error())
	}
	return bookObj, errorMsg
}

func GetAvailableBooks(userid string) (models.ErrorMsg, []models.Book) {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var books []models.Book
	var errorMsg models.ErrorMsg

	cursor, err := bookCollection.Find(ctx, bson.M{"userid": bson.D{{"$ne", userid}}, "availability": "YES"})

	if err != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.ERROR_FINDING_AVAILABLE_BOOKS
		logrus.WithFields(logrus.Fields{
			"Api":        "GetAvailableBooks",
			"StatusCode": 500,
		}).Error(err.Error())
		return errorMsg, books
	}

	defer cursor.Close(ctx)
	var book models.Book
	for cursor.Next(ctx) {
		if err = cursor.Decode(&book); err != nil {
			logrus.WithFields(logrus.Fields{
				"Api": "GetAvailableBooks",
			}).Error(err.Error())
		}

		books = append(books, book)
	}
	return errorMsg, books
}

func GetBooksByNameOrAuthorOrGenre(userid string, bookname string, authorname string, genre string) (models.ErrorMsg, []models.Book) {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var books []models.Book
	var userObj models.User
	var errorMsg models.ErrorMsg

	err := userCollection.FindOne(ctx, bson.M{"userid": userid}).Decode(&userObj)

	if err != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.USERID_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByNameOrAuthorOrGenre",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err.Error())
		return errorMsg, books
	}

	cursor, err1 := bookCollection.Find(ctx,
		bson.D{

			{"$or", []bson.M{{"userid": bson.D{{"$ne", userid}}}}},
			{"$or", []bson.M{{"bookname": bookname}, {"authorname": authorname}, {"genre": genre}}},
		},
	)

	if err1 != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.ERROR_FINDING_AVAILABLE_BOOKS
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByNameOrAuthorOrGenre",
			"StatusCode": 500,
		}).Error(err.Error())
		return errorMsg, books
	}

	defer cursor.Close(ctx)
	var book models.Book
	for cursor.Next(ctx) {
		if err = cursor.Decode(&book); err != nil {
			logrus.WithFields(logrus.Fields{
				"Api": "GetAvailableBooks",
			}).Error(err.Error())
		}
		fmt.Println(book)
		books = append(books, book)
	}
	return errorMsg, books
}

func GetBooksByUserName(userid string, lendername string) (models.ErrorMsg, []models.Book) {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var books []models.Book
	var userObj models.User
	var errorMsg models.ErrorMsg

	err := userCollection.FindOne(ctx, bson.M{"userid": userid}).Decode(&userObj)

	if err != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.USERID_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByUserName",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err.Error())
		return errorMsg, books
	}

	var err1 error
	var cursor *mongo.Cursor

	if lendername == "" {
		cursor, err1 = bookCollection.Find(ctx, bson.M{"userid": bson.M{"$ne": userid}})
	} else {
		cursor, err1 = bookCollection.Find(ctx,
			bson.D{

				{"$and", []bson.M{{"userid": bson.D{{"$ne", userid}}}, {"lendername": lendername}}},
			},
		)
	}

	if err1 != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.ERROR_FINDING_AVAILABLE_BOOKS
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByUserName",
			"StatusCode": 500,
		}).Error(err.Error())
		return errorMsg, books
	}

	defer cursor.Close(ctx)
	var book models.Book
	for cursor.Next(ctx) {
		if err = cursor.Decode(&book); err != nil {
			logrus.WithFields(logrus.Fields{
				"Api": "GetBooksByUserName",
			}).Error(err.Error())
		}
		fmt.Println(book)
		books = append(books, book)
	}
	return errorMsg, books
}

func BorrowBook(userid string, bookname string, lendername string, bookrating float64) models.ErrorMsg {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var errorMsg models.ErrorMsg
	var userObj1 models.User
	var userObj2 models.User
	var bookObj models.Book

	err := userCollection.FindOne(ctx, bson.M{"userid": userid}).Decode(&userObj1)

	if err != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.USERID_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "BorrowBook",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err.Error())
		return errorMsg
	}

	userObj1.Token = userObj1.Token - 1
	userCollection.FindOneAndReplace(ctx, bson.M{"userid": userid}, userObj1)

	filter := bson.D{{"$and", []bson.M{{"userid": bson.D{{"$ne", userid}}}, {"bookname": bookname}, {"lendername": lendername}, {"bookrating": bookrating}}}}

	err1 := bookCollection.FindOne(ctx, filter).Decode(&bookObj)
	if err1 != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.BOOK_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "BorrowBook",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err1.Error())
		return errorMsg
	}

	lenderId := bookObj.UserId

	err2 := userCollection.FindOne(ctx, bson.M{"userid": lenderId}).Decode(&userObj2)
	fmt.Println(err2)

	userObj2.Token = userObj2.Token + 1
	userCollection.FindOneAndReplace(ctx, bson.M{"userid": lenderId}, userObj2)

	if bookObj.Availability == "NO" {
		errorMsg.Code = 403
		errorMsg.Message = utils.BOOK_NOT_AVAILABLE
	} else {
		_, err3 := bookCollection.UpdateOne(
			ctx, filter,
			bson.D{
				{"$set", bson.D{{"availability", "NO"}, {"borrowedby", userid}}},
			},
		)

		if err3 != nil {
			errorMsg.Code = 500
			errorMsg.Message = utils.BORROW_BOOK_FAILED
			logrus.WithFields(logrus.Fields{
				"Api":        "BorrowBook",
				"StatusCode": 500,
				"UserID":     userid,
			}).Error(err3.Error())

		}

		bookObj.Status = "BORROWED"
		bookObj.Availability = "NO"
		bookObj.BorrowedBy = userid

		userObj1.Books = append(userObj1.Books, bookObj)
		userCollection.FindOneAndReplace(ctx, bson.M{"userid": userid}, userObj1)

		for i := 0; i < len(userObj2.Books); i++ {
			if userObj2.Books[i].BookId == bookObj.BookId {
				userObj2.Books[i].Availability = "NO"
				userObj2.Books[i].BorrowedBy = userid
				userCollection.FindOneAndReplace(ctx, bson.M{"userid": lenderId}, userObj2)
			}
		}

	}

	return errorMsg
}

func GetBooksByStatus(userid string, status string) (models.ErrorMsg, []models.Book) {
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var books []models.Book
	var userObj models.User
	var errorMsg models.ErrorMsg

	err := userCollection.FindOne(ctx, bson.M{"userid": userid}).Decode(&userObj)

	if err != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.USERID_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "GetBooksByStatus",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err.Error())
		return errorMsg, books
	}

	if status == "" {
		books = userObj.Books
	} else if status == "ADDED" {
		for _, book := range userObj.Books {
			if book.Status == "ADDED" {
				books = append(books, book)
			}
		}
	} else if status == "BORROWED" {
		for _, book := range userObj.Books {
			if book.Status == "BORROWED" {
				books = append(books, book)
			}
		}
	} else if status == "RETURNED" {
		for _, book := range userObj.Books {
			if book.Status == "RETURNED" {
				books = append(books, book)
			}
		}
	} else {
		errorMsg.Code = 403
		errorMsg.Message = "Status can be only ADDED OR BORROWED OR RETURNED.Try again.."
	}

	return errorMsg, books

}

func ReturnBook(userid string, bookname string, lendername string, bookrating float64) models.ErrorMsg {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var errorMsg models.ErrorMsg
	var userObj1 models.User
	var userObj2 models.User
	var bookObj models.Book

	err := userCollection.FindOne(ctx, bson.M{"userid": userid}).Decode(&userObj1)

	if err != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.USERID_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "ReturnBook",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err.Error())
		return errorMsg
	}

	userObj1.Token = (userObj1.Token) + 1

	filter := bson.D{{"$and", []bson.M{{"userid": bson.D{{"$ne", userid}}}, {"bookname": bookname}, {"lendername": lendername}, {"bookrating": bookrating}}}}

	err1 := bookCollection.FindOne(ctx, filter).Decode(&bookObj)
	if err1 != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.BOOK_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "ReturnBook",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err1.Error())
		return errorMsg
	}

	lenderId := bookObj.UserId

	userCollection.FindOne(ctx, bson.M{"userid": lenderId}).Decode(&userObj2)

	userObj2.Token = (userObj2.Token) - 1

	_, err3 := bookCollection.UpdateOne(
		ctx, filter,
		bson.D{
			{"$set", bson.D{{"availability", "YES"}, {"borrowedby", ""}}},
		},
	)

	if err3 != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.RETURN_BOOK_FAILED
		logrus.WithFields(logrus.Fields{
			"Api":        "ReturnBook",
			"StatusCode": 500,
			"UserID":     userid,
		}).Error(err3.Error())

	}

	for i := 0; i < len(userObj1.Books); i++ {
		if userObj1.Books[i].BookId == bookObj.BookId {
			userObj1.Books[i].Status = "RETURNED"
			userObj1.Books[i].Availability = "YES"
			userCollection.FindOneAndReplace(ctx, bson.M{"userid": userid}, userObj1)
		}
	}

	for i := 0; i < len(userObj2.Books); i++ {
		if userObj2.Books[i].BookId == bookObj.BookId {

			userObj2.Books[i].Availability = "YES"
			userObj2.Books[i].BorrowedBy = ""
			userCollection.FindOneAndReplace(ctx, bson.M{"userid": lenderId}, userObj2)
		}
	}

	return errorMsg

}

func RemoveBook(userid string, bookid string) models.ErrorMsg {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var errorMsg models.ErrorMsg
	var bookObj models.Book

	// filter1 := bson.D{{"$and", []bson.M{{"userid": userid}, {"bookid": bookid}}}}

	filter1 := bson.M{"userid": userid, "bookid": bookid}

	err1 := bookCollection.FindOne(ctx, filter1).Decode(&bookObj)

	if err1 != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.BOOK_NOT_FOUND_TO_DELETE
		logrus.WithFields(logrus.Fields{
			"Api":        "RemoveBook",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err1.Error())
		return errorMsg
	}

	filter3 := bson.M{"userid": bookObj.BorrowedBy}
	statement1 := bson.M{"$pull": bson.M{"books": bson.M{"bookid": bookid}}}

	_, err4 := userCollection.UpdateOne(ctx, filter3, statement1)

	if err4 != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.REMOVE_BOOK_FAILED
		logrus.WithFields(logrus.Fields{
			"Api":        "RemoveBook",
			"StatusCode": 500,
			"UserID":     userid,
		}).Error(err1.Error())
	}

	_, err2 := bookCollection.DeleteOne(ctx, filter1)
	if err2 != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.REMOVE_BOOK_FAILED
		logrus.WithFields(logrus.Fields{
			"Api":        "RemoveBook",
			"StatusCode": 500,
			"UserID":     userid,
		}).Error(err2.Error())
	}

	filter2 := bson.M{"userid": userid}
	statement := bson.M{"$pull": bson.M{"books": bson.M{"bookid": bookid}}}

	_, err3 := userCollection.UpdateOne(ctx, filter2, statement)

	if err3 != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.REMOVE_BOOK_FAILED
		logrus.WithFields(logrus.Fields{
			"Api":        "RemoveBook",
			"StatusCode": 500,
			"UserID":     userid,
		}).Error(err1.Error())
	}

	return errorMsg
}

func RateBook(userid string, bookid string, bookrating float64) models.ErrorMsg {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var errorMsg models.ErrorMsg
	var userObj1 models.User
	var userObj2 models.User
	var bookObj models.Book

	err := userCollection.FindOne(ctx, bson.M{"userid": userid}).Decode(&userObj1)

	if err != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.USERID_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "RateBook",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err.Error())
		return errorMsg
	}

	var flag int = 0
	var i int
	for i = 0; i < len(userObj1.Books); i++ {
		if userObj1.Books[i].BookId == bookid && (userObj1.Books[i].Status == "BORROWED" || userObj1.Books[i].Status == "RETURNED") {
			flag = 1
			break
		}
	}
	if flag == 0 {
		errorMsg.Code = 409
		errorMsg.Message = "To rate book user should have borrowed or returned that book"
		return errorMsg
	}

	if bookrating > 5 {
		errorMsg.Code = 403
		errorMsg.Message = "Rating should be out of 5"
		return errorMsg
	}

	filter1 := bson.M{"bookid": bookid}

	err1 := bookCollection.FindOne(ctx, filter1).Decode(&bookObj)
	if err1 != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.BOOK_NOT_FOUND_TO_RATE
		logrus.WithFields(logrus.Fields{
			"Api":        "RateBook",
			"UserId":     userid,
			"StatusCode": 404,
		}).Error(err1.Error())
		return errorMsg
	}

	if bookObj.BookRating == 0 {
		bookObj.BookRating = bookrating
	} else {
		bookObj.BookRating = ((bookObj.BookRating + bookrating) / 2)
	}

	bookCollection.FindOneAndReplace(ctx, filter1, bookObj)

	userObj1.Books[i].BookRating = bookObj.BookRating
	userCollection.FindOneAndReplace(ctx, bson.M{"userid": userid}, userObj1)

	userCollection.FindOne(ctx, bson.M{"userid": bookObj.UserId}).Decode(&userObj2)
	for i := 0; i < len(userObj2.Books); i++ {
		if userObj2.Books[i].BookId == bookObj.BookId {
			userObj2.Books[i].BookRating = bookObj.BookRating
		}
		userCollection.FindOneAndReplace(ctx, bson.M{"userid": bookObj.UserId}, userObj2)
	}
	return errorMsg
}
