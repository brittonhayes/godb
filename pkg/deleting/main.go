package main

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

// One function deletes one entry from db
func One(db *scribble.Driver, collection string, entry string) {
	if err := db.Delete(collection, entry); err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Deleted:", collection, entry)
}

// All function deletes the collection from the db
func All(db *scribble.Driver, collection string) {
	if err := db.Delete(collection, ""); err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Deleted:", collection)
}
