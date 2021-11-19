package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var x string = "absgtttdtd"
	z := sha256.Sum256([]byte(x))
	y := string(z[:])
	fmt.Println(y)

}
