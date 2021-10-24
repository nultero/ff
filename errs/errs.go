package errs

import (
	"fmt"
	"os"
)

func quit(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func Throw(err string) {
	s := fmt.Sprintf("%s %s", redError(), err)
	quit(s)
}

func ThrowSys(err error) {
	s := fmt.Sprintf("%s %s", redError(), err)
	quit(s)
}

func ThrowDuplArgError(this, prevFound string) {
	fmt.Printf("! >> '%s' found, but already passed '%s' as argument \n", this, prevFound)
	fmt.Println(redError(), "cannot have two of the same type of argument")
	os.Exit(1)
}

func redError() string {
	return "\033[31;1;4merror:\033[0m"
}

func ThrowQuiet(s string) {
	fmt.Println(s)
	os.Exit(0)
}
