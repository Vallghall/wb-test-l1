package main

import (
	"fmt"
	"strings"
)

// Set is a partial implementation of a Set adt
type Set[T comparable] struct {
	size int            // set size
	um   map[T]struct{} // underlying map
}

func (s Set[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("Set(")
	for i := range s.um {
		sb.WriteString(fmt.Sprintf("%v,", i))
	}
	sb.WriteString(");")
	return strings.Replace(sb.String(), ",)", ")", 1)
}

// New is a constructor func for type set
func New[T comparable]() *Set[T] {
	return &Set[T]{
		size: 0,
		um:   make(map[T]struct{}),
	}
}

// From is a constructor func for type set
func From[T comparable](arr []T) *Set[T] {
	res := New[T]()
	for _, val := range arr {
		res.Add(val)
	}

	return res
}

// Add adds value to the Set
func (s *Set[T]) Add(value T) {
	if s.Has(value) {
		return
	}

	s.um[value] = struct{}{}
	s.size++
}

// Has returns true if the given value exists within
// the Set, or false if not
func (s Set[T]) Has(value T) bool {
	_, ok := s.um[value]
	return ok
}

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"} // initial data
	s := From(data)                                      // create Set from data
	fmt.Println(s)                                       // print out the Set
}
