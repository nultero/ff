package main

import (
	"ff/colors"
	"fmt"
	"math"
	"strings"
)

const maxCols = 4 // default
const padding = 2 // default

func listOpts(files []string) {
	files = trimExtensions(files)
	f := sliceToGrid(files)
	maxLens, absMax := maxStrLens(f)
	printOutLines(f, maxLens)
	printLastLines(absMax)
}

func sliceToGrid(files []string) [][]string {

	output := [][]string{}

	floatRows := float64(len(files) / maxCols)
	numRows := int(math.Round(floatRows))

	idx := 0
	for i := 0; i < numRows; i++ {
		strs := []string{}
		for x := 0; x < maxCols; x++ {
			strs = append(strs, files[idx])
			idx++
		}

		output = append(output, strs)
	}

	return output
}

func maxStrLens(rrws [][]string) ([]int, int) {

	m := []int{} // stores max lengths within each column
	for x := 0; x < maxCols; x++ {
		m = append(m, 0)
	}

	for _, rw := range rrws {
		for i, s := range rw {
			if len(s) > m[i] {
				m[i] = len(s)
			}
		}
	}

	absMax := 0
	for i := range m {
		m[i] += padding
		absMax += m[i]
	}

	return m, absMax
}

func printOutLines(rrws [][]string, maxLens []int) {
	for _, rw := range rrws {

		line := colors.Blue("|") + " "

		for i, s := range rw {
			spaces := maxLens[i] - len(s)
			line += s + strings.Repeat(" ", spaces)
		}

		fmt.Println(line)
	}
}

func printLastLines(absMax int) {
	ln := colors.Blue("|")
	blip := colors.Blue("|>")
	ff := colors.Emph("ff")
	dash := colors.Blue("-")
	sepLine := fmt.Sprintf("%s%s", ln, strings.Repeat(dash, absMax-1))

	fmt.Println(sepLine)
	fmt.Println(blip, "No args were passed to", ff)
	fmt.Println(blip, "opts listed above")
}

func trimExtensions(files []string) []string {
	for i := range files {
		if len(files[i]) > 4 {
			ln := len(files[i])
			files[i] = files[i][:ln-4] // trim `.txt`
		}
	}

	return files
}
