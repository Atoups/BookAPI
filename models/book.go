
// Struct for attributes of a book.
//
// Defines and represents a book object and it's attributes.

package models

import "gopkg.in/mgo.v2/bson"

// Defines and represents a Book object and it's attributes, attribute types, and struct(object) formatting.
type Book struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`                  //Id attribute - Unique Book Identification
	Title       string        `json:"title" bson:"title"`             //Title attribute - Title of the book
	Author      string        `json:"author" bson:"author"`           //Author attribute - Who wrote the book
	Publisher   string        `json:"publisher" bson:"publisher"`     //Publisher attribute - Who published the book
	PublishDate string        `json:"publishdate" bson:"publishdate"` //PublishDate attribute - Date book was published
	Rating      int           `json:"rating" bson:"rating"`           //Rating attribute - Holds the rating of the book
	Status      string        `json:"status" bson:"status"`           //Status attribute - Status of the book (In or Out)
}
