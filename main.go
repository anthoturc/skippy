package main

import (
	"fmt"

	"github.com/anthoturc/go-skippy/set"
)

func main() {
	set := set.NewSkipListSet()

	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	fmt.Println(set)
	set.Delete(1)
	fmt.Println(set)
	set.Delete(2)
	fmt.Println(set)
}
