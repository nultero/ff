package main

import (
	"ff/errs"
	"os"
)

func createFFDir() {
	dir := getFFDir()
	err := os.Mkdir(dir, 0777)
	if err != nil {
		errs.ThrowSys(err)
	}

	confPath := dir + "/.config"
	err = os.WriteFile(confPath, []byte(""), 0777)
	if err != nil {
		errs.ThrowSys(err)
	}
}
