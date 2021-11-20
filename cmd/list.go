package cmd

import (
	"fmt"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list [TAG]",
	Short:   "list the existing notes under a given tag, or all notes if not specified",
	Aliases: []string{"l"},

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd := tics.MakeT("list").Blue().Str()
			tics.ThrowTooFewArgs(cmd)
		}

		s := viper.GetString("notesPath")
		s = tics.MakeT(s).Blue().Bold().Str()
		fmt.Println(s)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
