package main

import (
	"fmt"

	"github.com/smallnest/sqids"
)

func main() {
	s, _ := sqids.New(sqids.WithBlocklist([]string{"86Rf07"}))
	id, _ := s.Encode(1, 2, 3) // "se8ojk"
	numbers := s.Decode(id)    // [1, 2, 3]

	fmt.Println(id, numbers)
}
