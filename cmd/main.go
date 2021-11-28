package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	var err1 error
	err1 = errors.Wrap(err1, "error1")
	err1 = errors.Wrap(err1, "error2")
	err1 = errors.Wrap(err1, "errro3")
	fmt.Printf("%s", err1.Error())
}
