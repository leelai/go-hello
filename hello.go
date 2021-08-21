package main

import (
	"fmt"
	"log"

	"example.com/user/hello/morestrings"
	"github.com/denisbrodbeck/machineid"
	"github.com/google/go-cmp/cmp"
)

func main() {
	// fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
