package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var underCmd = &cobra.Command{
	Use:   "under [tag name]",
	Short: "open a note file to a given tag",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		fmt.Println(arg)

	},
}

func init() {
	rootCmd.AddCommand(underCmd)
}
