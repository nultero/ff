package cmd

import (
	"ff/cmd/add"
	"fmt"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [FILE_NAME]",
	Aliases: []string{"a", "make"},
	Short:   "add a note tag to a file",

	Args: cobra.MaximumNArgs(1),

	// TODO completion func, just grab the index file and go by file names
	Run: addFunc,
}

func addFunc(cmd *cobra.Command, args []string) {
	checkFlags()

	if len(args) == 0 {
		args = append(args, add.GetArg(&CategoryFlag, &NoteFlag))
	}

	arg := args[0]

	if CategoryFlag {
		fmt.Println("placeholder for category")

	} else {

		jump(arg)

		dex.Changed = true

		// TODO implement the autofmt after some kind of jump
	}

}

func init() {

	categoryStr := "creates a new note category " + tics.Make("*(will start with no tags)").Green().String()

	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&CategoryFlag, "category", "c", false, categoryStr)
	addCmd.Flags().BoolVarP(&NoteFlag, "new", "n", false, "")
	addCmd.Flags().StringVarP(&FileFlag, "file", "f", "", "specifies a notes file to search in")
}
