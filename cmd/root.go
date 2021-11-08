package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var defaultFFconf string = "$USER/.config/.ff.yaml"
var defaultFFpath string = "$USER/.ff"

var defaultConfSettings string = fmt.Sprintf(
	"notesPath: %v"+
		"\nmaxColumns: 4"+
		"\ncolumnPadding: 2", defaultFFpath)

var rootCmd = &cobra.Command{
	Use:   "ff",
	Short: "a cli notes indexing & quick-write tool  \n \x1b[31;1mf's in the chat\x1b[0m",

	// When called with no args, ff will list out notes index
	// in columnized fmt, along with footer about supplying arg.
	Run: func(cmd *cobra.Command, args []string) {
		ColumnizeAllNotes(viper.GetString("notesPath"))
	},
}

// Strictly the variant for when no args are passed to ff.
func ColumnizeAllNotes(ffDir string) {
	files := getFFfiles(ffDir)
	maxCols, err := strconv.Atoi(viper.GetString("maxColumns"))
	if err != nil {
		tics.ThrowSys(err)
	}

	maxLens := getZeroedIntSlice(maxCols)
	rows := [][]string{}

	for len(files) != 0 {
		rw := []string{}
		for i := 0; i < maxCols; i++ {
			if len(files) == 0 {
				break
			}

			if len(files[0]) > maxLens[i] {
				maxLens[i] = len(files[0])
			}

			rw = append(rw, files[0])
			files = files[1:]
		}
		rows = append(rows, rw)
	}

	maxLens = addPadding(maxLens)
	for _, rw := range rows {
		ln := ""
		for i, s := range rw {
			p := strings.Repeat(" ", maxLens[i]-len(s))
			ln += s + p
		}
		fmt.Println(ln)
	}

}

func getFFfiles(dir string) []string {
	files := []string{}
	fInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		tics.ThrowSys(err)
	}

	for _, f := range fInfo {
		files = append(files, f.Name()[:len(f.Name())-4])
	}

	return files
}

func getZeroedIntSlice(maxLen int) []int {
	s := []int{}
	for i := 0; i < maxLen; i++ {
		s = append(s, 0)
	}
	return s
}

func addPadding(mx []int) []int {
	p, err := strconv.Atoi(viper.GetString("columnPadding"))
	if err != nil {
		tics.ThrowSys(err)
	}

	for i := range mx {
		mx[i] += p
	}

	return mx
}

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
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		defaultConfSettings = strings.ReplaceAll(defaultConfSettings, "$USER", home)
		path := strings.ReplaceAll(defaultFFconf, "$USER", home)
		tics.MakeConfigPrompt(path, "ff", defaultConfSettings)

		path = strings.ReplaceAll(defaultFFpath, "$USER", home)
		tics.MakeDirPrompt(path, "ff")
		return
	}
}
