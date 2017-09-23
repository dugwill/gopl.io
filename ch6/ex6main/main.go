package main

import (
	"fmt"
	"github.com/dugwill/gopl.io/ch6/ex6intset"
)

func main() {

	var x, y ex6intset.IntSet
	ex6intset.x.Add(1)
	ex6intset.x.Add(144)
	ex6intset.x.Add(9)

	fmt.Println(ex6intset.x.String())

	ex6intset.y.Add(9)
	ex6intset.y.Add(42)
	fmt.Println(ex6intset.y.String())

}
