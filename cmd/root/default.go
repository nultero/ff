package root

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/nultero/tics"
	"github.com/spf13/viper"
)

// Strictly the variant for when no args are passed to ff.
func DefaultColumnize(ffDir string) {

	files := getFFfiles(ffDir)
	maxCols, err := strconv.Atoi(viper.GetString("max columns"))
	if err != nil {
		tics.ThrowSys(DefaultColumnize, err)
	}

	maxLens := make([]int, maxCols)
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
	dashStr := getDashesStr(&maxLens)
	fmt.Println(" \n" + dashStr) //buffer newln and dashes

	for _, rw := range rows {
		ln := ""
		for i, s := range rw {
			p := strings.Repeat(" ", maxLens[i]-len(s))
			ln += s + p
		}
		fmt.Println(ln)
	}

	s := []string{
		dashStr,
		tics.Make("+ no args passed to ff").Blue().String(),
		tics.Make("+ opts listed above").Blue().String(),
	}

	for _, v := range s {
		fmt.Println(v)
	}
}

func getFFfiles(dir string) []string {
	files := []string{}
	fInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		tics.ThrowSys(getFFfiles, err)
	}

	for _, f := range fInfo {
		files = append(files, f.Name()[:len(f.Name())-4])
	}

	return files
}

func addPadding(mx []int) []int {
	p, err := strconv.Atoi(viper.GetString("column padding"))
	if err != nil {
		tics.ThrowSys(addPadding, err)
	}

	for i := range mx {
		mx[i] += p
	}

	return mx
}

// Returns a string of dashes as long as the longest
// of the columnized chunks; I think it looks good as a TUI.
func getDashesStr(maxLens *[]int) string {
	numDashes := 0
	for _, i := range *maxLens {
		numDashes += i
	}

	return tics.Make("-").Repeat(numDashes).String()
}
