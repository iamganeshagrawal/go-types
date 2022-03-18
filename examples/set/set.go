package main

import (
	"fmt"

	"github.com/iamganeshagrawal/go-types/pkg/sets"
)

func main() {
	s := sets.NewSet[int]()
	s.Add(1, 2, 3, 4, 4)
	fmt.Println(s)
	fmt.Println(s.Size())
	s.Clear()
	fmt.Println(s)
	fmt.Println(s.Size())

	a := sets.NewSet(1, 2)
	b := sets.NewSet(2, 3)
	fmt.Println(a, b)
	c := a.Union(b)
	fmt.Println(a, b)
	fmt.Println(c)
	d := a.Intersect(b)
	fmt.Println(a, b)
	fmt.Println(d)
}
