

// Main application file to define mux and run application.
//
// Starts up the book api application to perform CRUD operations to manage a list of Books.
// Connects to the MongoDB session amd sets up the HTTP request routes using the gorilla/mux library.
// After defining the routes and the associated methods, the book api application runs and listens on port 8080.

package main

import (
	"github.com/gorilla/mux"
	"BookAPI/controllers"
	dao1 "BookAPI/dao"
	"log"
	"net/http"
)

// constant definitions
const (
	AppPortMessage 	           = ":8080"
	AppStartUp                 = "Application Started"
	BookEndpoint               = "/book"
	BookIdEndpoint             = "/book/{id}"
	ContentType                = "Content-Type"
	DBName                     = "BookMongo"
	DBServerUrl                = "mongodb://book_mongodb_1:27017"
	ListeningOnPort            = ("Listening on port ")
)

// BookDAO instance with MongoDB URL and 'BookMongo' Database
var dao = dao1.BookDAO{Server: DBServerUrl, Database: DBName}

// main function
//
// Defines HTTP request routes, connects to mongodb session and runs application server on port 8080
func main() {
	log.Println(AppStartUp)
	// connect to MongoDB through the dao instance
	dao.Connect()
	r := mux.NewRouter()

	// GET
	r.HandleFunc(BookEndpoint, controllers.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc(BookIdEndpoint, controllers.GetBook).Methods(http.MethodGet)

	// POST
	r.HandleFunc(BookEndpoint, controllers.CreateBook).Methods(http.MethodPost)

	// PUT
	r.HandleFunc(BookIdEndpoint, controllers.UpdateBook).Methods(http.MethodPut)

	// DELETE
	r.HandleFunc(BookIdEndpoint, controllers.DeleteBook).Methods(http.MethodDelete)

	// run application and listen on port 8080
	log.Println(ListeningOnPort, AppPortMessage)
	if err := http.ListenAndServe(AppPortMessage , r); err != nil {
		log.Fatal(err)
	}
}