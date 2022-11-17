package services

import (
	"encoding/json"
	"librarySystem/business"
	"librarySystem/interfaces"
	"librarySystem/models"
	"net/http"
	"strconv"

	"librarySystem/utils"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	interfaces.IUserService
}

func (service UserService) CreateUser(response http.ResponseWriter, request *http.Request) {

	var requestBody models.User
	json.NewDecoder(request.Body).Decode(&requestBody)

	logrus.WithFields(logrus.Fields{
		"Api":        "CreateUser",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.CREATE_USER_REQUEST)

	var resp models.ResponseUserObject
	var errorBody models.ErrorBody

	url, res := business.UserManager{}.CreateUser(requestBody)
	if url.Message != "" {
		errorBody.Error = url
		logrus.WithFields(logrus.Fields{
			"Api":        "CreateUser",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.CREATE_USER_FAILED)
		response.WriteHeader(url.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "CreateUser",
		"Path":       request.URL.Path,
		"StatusCode": 201,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.CREATE_USER_SUCCESS)

	resp.Data = res
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(resp)

	return

}

func (service UserService) RateUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userid := params["userid"]
	username := params["username"]
	userrating := params["userrating"]
	rating, _ := strconv.ParseFloat(userrating, 32)

	logrus.WithFields(logrus.Fields{
		"Api":        "RateUser",
		"Path":       request.URL.Path,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.RATE_USER_REQUEST)

	var msg models.ResponseDataString
	var resp models.ResponseDataSuccess
	var errorBody models.ErrorBody

	res := business.UserManager{}.RateUser(userid, username, rating)
	if res.Message != "" {
		errorBody.Error = res
		logrus.WithFields(logrus.Fields{
			"Api":        "RateUser",
			"Path":       request.URL.Path,
			"ClientAddr": request.RemoteAddr,
		}).Error(utils.RATE_USER_FAILED)
		response.WriteHeader(res.Code)
		json.NewEncoder(response).Encode(errorBody)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Api":        "RateUser",
		"Path":       request.URL.Path,
		"ErrorCode":  200,
		"UserId":     userid,
		"ClientAddr": request.RemoteAddr,
	}).Info(utils.RATE_USER_SUCCESS)
	msg.SuccessMsg = utils.SUCCESS_RATE_USER
	resp.Data = msg
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp)
}
