package main

import (
	"ff/cmd"
	"ff/errs"
	"fmt"
	"os"
)

const PATH = ".ff"

func main() {

	ff := dotfiles(PATH)
	args := os.Args[1:]

	if len(args) == 0 {
		listOpts(ff)

	} else {
		cmd.ParseArgs(args, ff)
	}

}

func getFFDir() string {
	hd, err := os.UserHomeDir()
	if err != nil {
		errs.ThrowSys(err)
	}

	return fmt.Sprintf("%s/%s", hd, PATH)
}

func dotfiles(path string) []string {

	dir := getFFDir()

	entries, err := os.ReadDir(dir)
	if err != nil {
		createFFDir()
	}

	files := []string{}
	for _, f := range entries {

		if f.Name()[0] == '.' {
			continue
		}

		files = append(files, f.Name())
	}

	return files
}
