package adding

import (
	"fmt"

	"github.com/brittonhayes/godb/pkg/types"

	scribble "github.com/nanobox-io/golang-scribble"
)

// Single function adds one entry to the database
func Single(db *scribble.Driver, folder string, file string, entry *types.Fish) {
	// Write a fish to the database
	if err := db.Write(folder, file, entry); err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("adding.Single has run on:", folder, file, entry)
}
