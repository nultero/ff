package cmd

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/nultero/tics"
	"github.com/spf13/viper"
)

// Strictly the variant for when no args are passed to ff.
func columnizeAllNotes(ffDir string) {
	fmt.Println("")

	files := getFFfiles(ffDir)
	maxCols, err := strconv.Atoi(viper.GetString("max columns"))
	if err != nil {
		tics.ThrowSysDescriptor(tics.BlameFunc(columnizeAllNotes), err)
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

	nDashes := 0
	for _, i := range maxLens {
		nDashes += i
	}

	s := []string{
		tics.MakeT("-").Blue().Repeat(nDashes).Str(),
		tics.MakeT("+ no args passed to ff").Blue().Str(),
		tics.MakeT("+ opts listed above").Blue().Str(),
	}

	for _, v := range s {
		fmt.Println(v)
	}
}

func getFFfiles(dir string) []string {
	files := []string{}
	fInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		tics.ThrowSysDescriptor(tics.BlameFunc(getFFfiles), err)
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
	p, err := strconv.Atoi(viper.GetString("column padding"))
	if err != nil {
		tics.ThrowSysDescriptor(tics.BlameFunc(addPadding), err)
	}

	for i := range mx {
		mx[i] += p
	}

	return mx
}
