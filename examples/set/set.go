package main

import (
	"fmt"

	"github.com/iamganeshagrawal/go-types/pkg/sets"
)

func main() {
	set := sets.NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println("Set Size:", set.Size())
	fmt.Println("Set:", set)
	fmt.Println("have 2:", set.Has(2))
	fmt.Println("have 13:", set.Has(13))
	fmt.Println("add 13:", set.AddIfNotExist(13))
	fmt.Println("have 13:", set.Has(13))
	fmt.Println("Set Size:", set.Size())
	fmt.Println("Set:", set)
	fmt.Println("add 13:", set.AddIfNotExist(13))
}
