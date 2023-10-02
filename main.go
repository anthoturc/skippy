package main

import (
	"fmt"

	"github.com/anthoturc/go-skippy/set"
)

func main() {
	set := set.NewSkipListSet()

	for i := 0; i < 10; i++ {
		set.Insert(i)
	}

	fmt.Println(set)
}
