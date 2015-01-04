package main

import (
	"fmt"
	rn "github.com/seiyria/restricted-number-go/restrictednumber"
)

func main() {
	fmt.Printf("Test")
	hp := rn.New()
	fmt.Printf("%v", hp.Val())
}
