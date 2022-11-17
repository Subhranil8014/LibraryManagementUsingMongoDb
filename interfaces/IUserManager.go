package interfaces

import "librarySystem/models"

type IUserManager interface {
	CreateUser(userObj models.User) (models.ErrorMsg, models.User)
	RateUser(userid string, username string, userrating float64) models.ErrorMsg
}
