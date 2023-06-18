package main

import (
	"fmt"
	"github.com/raspi/ranger"
)

func main() {
	r := []int8{-1, 0, 1, 3, 4, 5, 10, 11, 12, 13, 14, 15, 20, 30, 21, 25, 24, 21}

	ranges, err := ranger.GetIntegerRanges(r)
	if err != nil {
		panic(err)
	}

	for _, r := range ranges {
		fmt.Printf(`%v - %v`+"\n", r.Start, r.End)
	}
}
