package cmd

import (
	// "github.com/brittonhayes/godb/pkg/adding"
	// "github.com/brittonhayes/godb/pkg/reading"

	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/cobra"
)

func initialize(dir string) *scribble.Driver {
	// Initialize the database
	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	return db
}

// Execute runs the cli
func Execute() {
	db := initialize("./db")
	var (
		rootCmd = &cobra.Command{
			Use:   "godb",
			Short: "A local golang json db",
		}
	)
	rootCmd.AddCommand(Read(db))
	rootCmd.AddCommand(Write(db))
	rootCmd.Execute()
}
