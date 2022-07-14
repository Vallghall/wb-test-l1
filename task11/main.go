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

// Intersect method returns a new Set which is an
// intersection between two sets
func (s Set[T]) Intersect(another *Set[T]) *Set[T] {
	newSet := New[T]()

	for this := range s.um {
		for that := range another.um {
			if this == that {
				newSet.Add(this)
			}
		}
	}

	return newSet
}

func main() {
	a := []int{1, 2, 3, 4, 5} // init data for the first Set
	b := []int{3, 4, 5, 6, 7} // init data for the second Set

	s1, s2 := From(a), From(b) // init two Sets

	s3 := s1.Intersect(s2) // find the intersection between two sets

	fmt.Println(s3)
}
