package cmd

import (
	"fmt"

	"github.com/brittonhayes/godb/pkg/reading"
	"github.com/brittonhayes/godb/pkg/types"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/cobra"
)

// Read function reads an entry from the db
func Read(db *scribble.Driver) *cobra.Command {
	var (
		readCmd = &cobra.Command{
			Use:   "read [collection] [name] [species]",
			Short: "Read an entry from the database",
			Args:  cobra.MinimumNArgs(3),
			Run: func(cmd *cobra.Command, args []string) {
				//* Read a single fish from the database
				entry := types.Fish{Name: args[1], Species: args[2]}
				reading.One(db, args[0], args[1], &entry)
				fmt.Println(entry)
			},
		}
	)

	return readCmd
}
