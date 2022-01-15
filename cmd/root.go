package cmd

import (
	"errors"
	"ff/cmd/root"
	"fmt"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// When called with no args, ff will list out notes index
// in columnized fmt, along with footer about supplying arg.
var rootCmd = &cobra.Command{
	Use:   "ff",
	Short: "a cli notes indexing & quick-write tool  \n" + flavorText,

	// TODOO use index instead of dir calls in root columnizer
	Run: func(cmd *cobra.Command, args []string) {
		s := viper.GetString("notes path")
		root.DefaultColumnize(s)
	},
}

// A primitive in root.go to separate out some mutually exclusive flags.
// Doesn't necessarily need flagsets on just two or three flags.
func checkFlags() {
	if CategoryFlag && NoteFlag {
		s := fmt.Sprintf(
			"'%v' and '%v' are mutually exclusive flags",
			tics.Make("-c").Blue().String(),
			tics.Make("-n").Blue().String(),
		)
		err := errors.New(s)
		tics.ThrowSys(checkFlags, err)
	}
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())

	if dex.Changed { // only some cmds will have triggered this,
		dex.Index() // prompts re-indexing and autofmt
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	confMap = tics.CobraRootInitBoilerPlate(confMap, true)
	confPath := confMap[confFile]
	viper.SetConfigFile(confPath)
	viper.AutomaticEnv()

	// If a config file is found, read it in, else make one with prompt.
	err := viper.ReadInConfig()
	if err != nil {
		tics.RunConfPrompts("ff", confMap, defaultSettings)
		tics.ThrowQuiet("")
	}
}
