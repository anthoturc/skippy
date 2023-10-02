package main

import (
	"fmt"

	"github.com/anthoturc/go-skippy/list"
)

func main() {
	skipList := list.NewSkipList()
	for i := 9; i >= 1; i -= 1 {
		skipList.Insert(i)
	}
	fmt.Println(skipList)
}
