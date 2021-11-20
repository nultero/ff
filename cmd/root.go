package cmd

import (
	"fmt"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var flavorText = tics.Red(" -> f's in the chat")

var confMap = map[string]string{
	"confDir": "$USER/.config/",
	"dataDir": "$USER/.ff",
}

var defaultSettings = []string{
	"notes path: $USER/.ff",
	"max columns: 4",
	"column padding: 2",
}

var rootCmd = &cobra.Command{
	Use:   "ff",
	Short: "a cli notes indexing & quick-write tool  \n" + flavorText,

	// When called with no args, ff will list out notes index
	// in columnized fmt, along with footer about supplying arg.
	Run: func(cmd *cobra.Command, args []string) {
		s := viper.GetString("notes path")
		fmt.Println("viper notes path:", s)
		columnizeAllNotes(s)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/ff.yaml)")
}

func initConfig() {

	confMap = tics.CobraRootInitBoilerPlate(confMap, true)
	confPath := confMap["confDir"]

	viper.AddConfigPath(confPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("ff")
	viper.AutomaticEnv()

	// If a config file is found, read it in, else make one with prompt.
	err := viper.ReadInConfig()
	if err != nil {
		tics.RunConfPrompts("ff", confMap, defaultSettings)
		return
	}
}
