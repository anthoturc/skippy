package main

import (
	"fmt"

	"github.com/anthoturc/go-skippy/set"
)

func main() {
	skipList := set.NewSkipListSet()
	for i := 9; i >= 1; i -= 1 {
		skipList.Insert(i)
	}
	fmt.Println(skipList)
}
