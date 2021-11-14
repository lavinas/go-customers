package main

import (
	"fmt"

	"github.com/dongri/phonenumber"
)

func main() {
	number := phonenumber.GetISO3166ByNumber("447400123456", false)
	fmt.Println(number)
}
