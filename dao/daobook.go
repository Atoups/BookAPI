

// Data Access Object (DAO) struct implementation to MongoDB.
//
// Handles the MongoDB database CRUD operations such as Find, FindId, Insert, Delete, and UpdateId.
// The DAO is also responsible for connecting to the MongoDB session, or Closing the session with the
// Connect() and Close() methods.

package dao

import (
	"BookAPI/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// A BookDAO represents the information needed to access MongoDB
type BookDAO struct {
	Server   string
	Database string
}

// variable referencing the mongo database as a type
var db *mgo.Database

// constant definitions
const (
	MsgClosedDB           = "Closed connection to MongoDB"
	Collection            = "Books"
	MsgConnDB             = "Connected to MongoDB"
)

// Connect
//
// The function establishes a connection session to MongoDB for the DAO.
func (b *BookDAO) Connect() {
	session, err := mgo.Dial(b.Server)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(MsgConnDB)
	db = session.DB(b.Database)
}

// Close
//
// The function closes the MongoDB session for the DAO.
func (b *BookDAO) Close() {
	db.Session.Close()
	log.Println(MsgClosedDB)
}

// FindAll
//
// Function that finds a list of all the books from the collection.
// The functions returns a slice of books and an error, which will be nil if the books are retrieved.
func (b *BookDAO) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := db.C(Collection).Find(nil).All(&books)
	return books, err
}

// FindByID
//
// Function that finds a book by its id from the collection for a given ID parameter.
// Function that returns a book struct and an error, which will be nil if the book is retrieved.
func (b *BookDAO) FindById(id string) (models.Book, error) {
	var book models.Book
	err := db.C(Collection).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

// Insert
//
// Function that inserts a new book from the collection for a given Book parameter.
// Function that return an error, which will be nil for a successful transaction or an error for any encountered errors.
func (b *BookDAO) Insert(book models.Book) error {
	err := db.C(Collection).Insert(&book)
	return err
}

// Delete function
//
// Function that deletes an existing book from the collection for a given Book parameter.
// Function that return an error, which will be nil for a successful transaction or an error for any encountered errors.
func (b *BookDAO) Delete(book models.Book) error {
	err := db.C(Collection).Remove(&book)
	return err
}

// Update function
//
// Function that updates an existing book from the collection for a given Book parameter.
// Function that returns an error, which will be nil for a successful transaction or an error for any encountered errors.
func (b *BookDAO) Update(book models.Book) error {
	err := db.C(Collection).UpdateId(book.Id, &book)
	return err
}