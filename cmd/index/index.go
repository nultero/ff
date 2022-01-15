package index

import (
	"os"
	"strings"

	"github.com/nultero/tics"
)

type Log struct {
	Changed   bool
	File      string
	DeltaFile string
	Indices   []string
}

// Re-orders ff's index dotfile,
// and if changes were made to any notes or
// categories, then it will re-store them for later.
func (l *Log) Index() {

}

// Pulls ff's index file into memory so that only 1 syscall is made
// to quickly parse and find any given note / tag.
func (l *Log) Read() {
	b, err := os.ReadFile(l.File)
	if err != nil {
		tics.ThrowSys(l.Read, err)
	}

	s := string(b)
	l.Indices = strings.Split(s, "\n")
}
