package cmd

import (
	"github.com/brittonhayes/godb/pkg/adding"
	"github.com/brittonhayes/godb/pkg/types"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/cobra"
)

// Write function adds an entry to the db
func Write(db *scribble.Driver) *cobra.Command {
	var (
		writeCmd = &cobra.Command{
			Use:   "write [collection] [name] [species]",
			Short: "Write an entry to the database",
			Args:  cobra.MinimumNArgs(3),
			Run: func(cmd *cobra.Command, args []string) {
				entry := types.Fish{Name: args[1], Species: args[2]}
				adding.Single(db, args[0], args[1], &entry)
			},
		}
	)

	return writeCmd
}
