package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var defaultFFconf string = "$USER/.config/.ff.yaml"
var defaultFFpath string = "$USER/.ff"

var defaultConfSettings string = fmt.Sprintf("notesPath: %v", defaultFFpath)

var rootCmd = &cobra.Command{
	Use:   "ff",
	Short: "a cli notes indexing & quick-write tool  \n \x1b[31;1mf's in the chat\x1b[0m",

	// When called with no args, ff will list out notes index
	// in columnized fmt, along with footer about supplying arg.
	Run: func(cmd *cobra.Command, args []string) {

		// Columnize func will go here after dir is created
		fmt.Println(strings.Repeat("/", 10), "called with no args, func not written yet")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.ff.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)

	} else {
		path, err := os.UserHomeDir()
		cobra.CheckErr(err)
		path += "/.config"

		viper.AddConfigPath(path)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ff")
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in, else make one with prompt.
	err := viper.ReadInConfig()
	if err != nil {
		path, err := os.UserHomeDir()
		cobra.CheckErr(err)

		defaultConfSettings = strings.ReplaceAll(defaultConfSettings, "$USER", path)
		path = strings.ReplaceAll(defaultFFconf, "$USER", path)
		tics.MakeConfigPrompt(path, "ff", defaultConfSettings)
	}
}
