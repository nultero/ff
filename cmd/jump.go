package cmd

import (
	"errors"
	"os"
	"os/exec"

	"github.com/nultero/tics"
	"github.com/spf13/viper"
)

// TODOOOO impl an index so that the whole directory doesn't have to
// be read into memory to impl the jumps
// TODOOO impl the jump for a tagname
func jump(arg string) {
	ed := viper.GetString("editor")
	checkEditor(ed)

	c := exec.Command(ed, "+50", "main.go") // prototype behavior
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		tics.ThrowSys(jump, err)
	}
}

// For my own sanity, if for some reason I use this on another machine
// and don't unpack my dotfiles properly, or I misconfigure something,
// this should come up as the obvious error.
func checkEditor(ed string) {
	if len(ed) == 0 {
		err := errors.New("no valid config / no editor specified to open notes with")
		tics.ThrowSys(checkEditor, err)
	}
}
