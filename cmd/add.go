package cmd

import (
	"fmt"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [FILE_NAME, SUB_HEADER_NAME]",
	Aliases: []string{"a", "make"},
	Short:   "Create and open a new file /  with the name listed.",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 2 {
			cmd := tics.MakeT("add").Blue().Str()
			tics.ThrowTooManyArgs(cmd)

		} else {
			fmt.Println("yes")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
