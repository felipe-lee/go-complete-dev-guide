/*
set is a little experiment of mine with implementing the set type in go, similar to set in python. It is definitely not
a thorough implementation. Mainly an exercise for myself. I did find this package which seems much better (though
doesn't seem actively maintained): https://github.com/deckarep/golang-set
 */
package main

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func NewSet(items []string) Set {
	var newSet = Set{}

	for _, item := range items {
		newSet.add(item)
	}

	return newSet
}

func (s Set) add(item string) {
	s[item] = true
}

func (s Set) remove(item string) {
	delete(s, item)
}

func (s Set) contains(item string) bool {
	_, itemPresent := s[item]
	return itemPresent
}

func (s Set) String() string {
	keys := make([]string, len(s))
	i := 0

	for key := range s {
		keys[i] = key
		i++
	}

	return strings.Join(keys, ", ")
}

func main() {
	var a = Set{}

	a.add("hi")
	a.add("bye")

	fmt.Println(a)

	a.remove("bye")

	fmt.Printf("%v\n", a)

	fmt.Println("is bye in a?", a.contains("bye"))

	b := NewSet([]string{"a", "b", "c"})

	fmt.Println(b)
}
