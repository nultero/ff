package cmd

import "fmt"

func ParseArgs(args, files []string) {
	for _, r := range args {
		fmt.Println(r)
	}
}
