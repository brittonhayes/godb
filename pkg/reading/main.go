package reading

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brittonhayes/godb/pkg/types"

	scribble "github.com/nanobox-io/golang-scribble"
)

// One reads an entry from the database
func One(db *scribble.Driver, folder string, file string, entry *types.Fish) error {
	result := entry
	response := db.Read(folder, file, &result)
	if response != nil {
		log.Fatal("Error", response)
	}

	fmt.Printf("\nREAD:\n\n")
	fmt.Printf("	Collection: %v\n", folder)
	fmt.Printf("	Object: %v\n", file)
	fmt.Printf("	Data: %+v\n", entry)

	return response
}

// All reads all entries from the database, unmarshaling the response.
func All(db *scribble.Driver) []types.Fish {
	records, err := db.ReadAll("fish")
	if err != nil {
		log.Fatal("Error", err)
	}

	results := []types.Fish{}
	for _, f := range records {
		resultFound := types.Fish{}
		if err := json.Unmarshal([]byte(f), &resultFound); err != nil {
			fmt.Println("Error", err)
		}
		results = append(results, resultFound)

		fmt.Printf("\nREAD:\n\n")
		fmt.Printf("	DATA: %v\n", resultFound)
	}
	return results
}
