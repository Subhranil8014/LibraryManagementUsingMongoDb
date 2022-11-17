package interfaces

import (
	"net/http"
)

type IUserService interface {
	CreateUser(response http.ResponseWriter, request *http.Request)
	RateUser(response http.ResponseWriter, request *http.Request)
}
