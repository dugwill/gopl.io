// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 167.

// Exercise 6.1: Instructions were to add the following Methods
// func (*IntSet) Len() int
// func (*IntSet) Remove(x int)
// func (*IntSet) Clear()
// func (*IntSet) Copy() *IntSet
// For this exercise I added a main func and changed the package name

// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

func main() {

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(66)

	fmt.Println("x: " + x.String())

	y.Add(9)
	y.Add(42)
	y.Add(291)
	fmt.Println("y: " + y.String())

	fmt.Printf("X Length: %v\n", x.Len())

	fmt.Println("x: " + x.String())
	fmt.Println("Removing 1 from x\n")
	x.Remove(1)
	fmt.Println(x.String())
	fmt.Println("Removing 14 from x\n")
	x.Remove(14)
	fmt.Println("x: " + x.String())

	fmt.Println("Clearing X")
	x.Clear()
	fmt.Println("X: " + x.String())
	fmt.Printf("X Length: %v\n\n", x.Len())

	fmt.Println("Duplicating y to z")
	z := y.Copy()
	fmt.Printf("Y: %v\n", y.String())
	fmt.Printf("Z: %v\n", z.String())

	for _, i := range z.Elems() {
		fmt.Printf("z element: %v\n", i)
	}

	x.Add(3)
	x.Add(144)
	x.Add(67)
	x.Add(123)

	fmt.Println("x:" + x.String())
	fmt.Println("y:" + y.String())
	x.UnionWith(&y)
	fmt.Println("Union of x and y:" + x.String())

	x.IntersectWith(&y)
	fmt.Println("Intersction of x and y:" + x.String())

	x.Add(69)
	y.Add(154)
	x.DifferenceWith(&y)
	fmt.Println("Difference between x and y" + x.String())

}

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// len returns the number of elements.
func (s *IntSet) Len() int {
	counter := 0
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				fmt.Printf("found bit! word:%v, bit:%v\n", i, j)
				counter++

			}
		}
	}
	return counter
}

// remove deletes the value x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	//for word >= len(s.words) {
	//	s.words = append(s.words, 0)
	//}

	if s.words[word]&(1<<uint(bit)) == 0 {
		fmt.Printf("%v not found\n", x)
	}

	s.words[word] ^= 1 << bit
	fmt.Printf("%v removed\n", x)
}

// clear deletes all elements from the set.
func (s *IntSet) Clear() {
	for word := range s.words {
		if s.words[word] != 0 {
			s.words[word] &= 0
		}
	}
}

// Copy duplicates the set
func (s *IntSet) Copy() *IntSet {

	var dup IntSet
	for word := range s.words {
		dup.words = append(dup.words, s.words[word])
	}

	return &dup
}


// Excercise 6.2 - Create a variadic function that adds a list of
// ints to the set
// AddAll takes a list of ints and, using the add func, adds them
// to the Set
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

//Exercise 6.3

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		s.words[i] &= tword
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Excercise 6.4
// Create func Elems that returns a slice suitable for ranging over
// Elems walks all the word in IntSet and appends the elements to
// a slice of ints, elements.
func (s *IntSet) Elems() []int {

	var elements []int

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				fmt.Printf("Appending %v\n", 64*i+j)
				elements = append(elements, 64*i+j)
			}
		}
	}
	return elements
}
