package main

import (
	"fmt"

	"github.com/anthoturc/skippy/set"
)

func main() {
	set := set.NewSkipListSet[int]()

	for i := 0; i < 10; i++ {
		set.Insert(i)
	}

	fmt.Println(set)
}
