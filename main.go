package main

import (
	"fmt"
	database "librarySystem/dbUtils"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")

func main() {
	router := initRouter()
	fmt.Println("Starting the server")
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3003"
	}
	err := http.ListenAndServe(port, handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))

	if err != nil {
		logrus.Error(err.Error())
	}
}
