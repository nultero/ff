package cmd

import (
	"ff/cmd/fsys"
	"fmt"
	"strings"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [TAG]",
	Short:   "list the existing note tags under a given section, or all note tags if not specified",
	Aliases: []string{"l"},

	ValidArgsFunction: getFFcompletions,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd := tics.Make("list").Blue().String() // TODO `list` without args just gets tags from index
			tics.ThrowTooFewArgs(cmd)                // don't keep as final behavior
		}

		arg := args[0]

		if dir, ok := confMap[dataDir]; ok {
			getFileMatches(arg, dir)
		}
	},
}

func getFileMatches(arg, dir string) {
	for _, f := range fsys.GetFFfilenames(dir) {
		if strings.Contains(f, arg) {
			bytes := tics.GetFile(dir + "/" + f + ".txt")
			f = strings.ReplaceAll(f, arg, tics.Make(arg).Blue().String())
			fmt.Println(f)
			fmt.Println(string(bytes))
			fmt.Printf("\n")
		}

	}
}

func getFFcompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	scd := cobra.ShellCompDirectiveNoFileComp

	if len(args) != 0 {
		return nil, scd
	}

	files := []string{}
	if d, ok := confMap[dataDir]; ok {
		files = fsys.GetFFfilenames(d)
	}

	return files, scd
}

func init() {
	rootCmd.AddCommand(listCmd)
}
