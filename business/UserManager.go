package business

import (
	"librarySystem/dao"
	"librarySystem/interfaces"
	"librarySystem/models"

	"github.com/sirupsen/logrus"
)

type UserManager struct {
	interfaces.IUserManager
}

func (manager UserManager) CreateUser(userObj models.User) (models.ErrorMsg, models.User) {
	res, err := dao.CreateUser(userObj)

	if err.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "CreateUser",
			"StatusCode": err.Code,
		}).Error(err.Message)
	}
	return err, res
}

func (manager UserManager) RateUser(userid string, username string, userrating float64) models.ErrorMsg {
	res := dao.RateUser(userid, username, userrating)

	if res.Message != "" {
		logrus.WithFields(logrus.Fields{
			"Api":        "RateUser",
			"StatusCode": res.Code,
			"UserId":     userid,
		}).Error(res.Message)

	}
	return res
}
