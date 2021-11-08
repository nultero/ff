package cmd

import (
	"fmt"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [TAG]",
	Short: "list the existing notes under a given tag, or all notes if not specified",
	Run: func(cmd *cobra.Command, args []string) {
		//
		s := tics.MakeT("CHAIN COLORS").Blue().Bold().Str()
		fmt.Println(s)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
