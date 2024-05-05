package main

import (
	"fmt"

	"github.com/smallnest/sqids"
)

func main() {
	s, _ := sqids.New(sqids.WithMinLength(10))
	id, _ := s.Encode(1, 2, 3) // "86Rf07xd4z"
	numbers := s.Decode(id)    // [1, 2, 3]

	fmt.Println(id, numbers)
}
