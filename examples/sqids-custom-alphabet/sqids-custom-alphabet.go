package main

import (
	"fmt"

	"github.com/smallnest/sqids"
)

func main() {
	s, _ := sqids.New(sqids.WithAlphabet("FxnXM1kBN6cuhsAvjW3Co7l2RePyY8DwaU04Tzt9fHQrqSVKdpimLGIJOgb5ZE"))
	id, _ := s.Encode(1, 2, 3) // "B4aajs"
	numbers := s.Decode(id)    // [1, 2, 3]

	fmt.Println(id, numbers)
}
