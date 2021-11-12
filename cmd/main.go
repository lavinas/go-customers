package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("YUU", "go")
	x := strings.ToUpper(os.Getenv("YUU"))
	if x == "" {
		fmt.Println("Blank")
	}
	fmt.Println(x)

	z := 19888222 % 100
	fmt.Println(z)

}
