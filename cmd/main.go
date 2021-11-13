package main

import (
	"fmt"

	"github.com/dongri/phonenumber"
)

func main() {
	number := phonenumber.ParseWithLandLine("11998987666", "BR")
	fmt.Println(number)
}
