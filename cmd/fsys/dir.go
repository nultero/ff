package fsys

import (
	"io/ioutil"

	"github.com/nultero/tics"
)

func GetFFfilenames(dir string) []string {
	files := []string{}
	fInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		tics.ThrowSys(GetFFfilenames, err)
	}

	for _, f := range fInfo {
		files = append(files, f.Name()[:len(f.Name())-4])
	}

	return files
}
