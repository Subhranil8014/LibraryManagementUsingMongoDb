package dao

import (
	"context"
	"fmt"
	database "librarySystem/dbUtils"
	"librarySystem/models"
	"librarySystem/utils"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(userObj models.User) (models.User, models.ErrorMsg) {
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
	var errorMsg models.ErrorMsg
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userObj.UserRating = 0
	userObj.UserId = generateUUID()
	loc, _ := time.LoadLocation("Asia/Kolkata")
	currentTime := time.Now().In(loc)
	t := currentTime.Format(time.RFC850)
	userObj.CreationTime = t

	_, err := userCollection.InsertOne(ctx, userObj)
	if err != nil {
		errorMsg.Code = 500
		errorMsg.Message = utils.CREATE_USER_FAILED
		logrus.WithFields(logrus.Fields{
			"Api":        "CreateUser",
			"StatusCode": 500,
			"ID":         userObj.UserId,
		}).Error(err.Error())
	}
	return userObj, errorMsg

}

func RateUser(userid string, username string, userrating float64) models.ErrorMsg {
	var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var errorMsg models.ErrorMsg
	var userObj models.User

	filter := bson.D{{"$and", []bson.M{{"userid": bson.D{{"$ne", userid}}}, {"username": username}}}}

	err1 := userCollection.FindOne(ctx, filter).Decode(&userObj)
	if err1 != nil {
		errorMsg.Code = 404
		errorMsg.Message = utils.USER_NOT_FOUND
		logrus.WithFields(logrus.Fields{
			"Api":        "RateUser",
			"StatusCode": 404,
		}).Error(err1.Error())
		return errorMsg
	}

	if userrating > 5 {
		errorMsg.Code = 403
		errorMsg.Message = "Rating should be out of 5"
		return errorMsg
	}
	if userObj.UserRating == 0 {
		userObj.UserRating = userrating
	} else {
		userObj.UserRating = ((userObj.UserRating + userrating) / 2)
	}

	userCollection.FindOneAndReplace(ctx, filter, userObj)

	filter1 := bson.D{{"$and", []bson.M{{"userid": bson.D{{"$ne", userid}}}, {"lendername": username}}}}

	_, err2 := bookCollection.UpdateOne(
		ctx, filter1,
		bson.D{
			{"$set", bson.D{{"lenderrating", userObj.UserRating}}}},
	)

	fmt.Println(err2)

	return errorMsg
}

func generateUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := uuidWithHyphen.String()

	return uuid
}
