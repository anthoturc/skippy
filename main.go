package main

import (
	"fmt"

	"github.com/anthoturc/go-skippy/list"
)

func main() {
	skipList := list.NewSkipList()
	skipList.Insert(1)
	// skipList.Insert(2)
	// skipList.Insert(3)
	fmt.Println(skipList)
}
