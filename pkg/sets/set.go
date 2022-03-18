package sets

import "fmt"

// NewSet returns a new Set with the given elements.
// Example:
//  setA := NewSet(1,2,3,4)
//  setB := NewSet("a", "b", "c")
//  setC := NewSet[int64](1,2,3)
func NewSet[T comparable](v ...T) *Set[T] {
	m := make(map[T]struct{}, len(v))
	for _, item := range v {
		m[item] = struct{}{}
	}
	return &Set[T]{
		values: m,
	}
}

// NewEmptySet returns a new empty Set.
// Example:
//  set := NewEmptySet[int]()
func NewEmptySet[T comparable]() *Set[T] {
	return &Set[T]{
		values: make(map[T]struct{}),
	}
}

// Set is an abstract data type that can store unique values, without any particular order.
// It is a computer implementation of the mathematical concept of a finite set
// and its implementation is using a Go map[T]struct{}.
type Set[T comparable] struct {
	values map[T]struct{}
}

// Size returns current size of the set.
func (s Set[T]) Size() int { return len(s.values) }

// Clear remove all the elements of the set container, thus making its size 0.
func (s *Set[T]) Clear() {
	s.values = make(map[T]struct{})
}

// Add will add an element or elements to the set.
func (s *Set[T]) Add(v ...T) {
	for _, item := range v {
		s.values[item] = struct{}{}
	}
}

// AddIfNotExist will add an element to the set, and return true if it was added
// or false if the value already existed in the set.
func (s *Set[T]) AddIfNotExist(e T) bool {
	if s.Has(e) {
		return false
	}
	s.values[e] = struct{}{}
	return true
}

// Has returns true if the element exists in the set.
func (s *Set[T]) Has(e T) bool {
	_, ok := s.values[e]
	return ok
}

// Contains returns true if all elements exists in the set.
func (s *Set[T]) Contains(v ...T) bool {
	for _, item := range v {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// Remove remove all elements from the set.
func (s *Set[T]) Remove(v ...T) {
	for _, item := range v {
		delete(s.values, item)
	}
}

// String returns string represention of the set.
func (s *Set[T]) String() string {
	return fmt.Sprint(s.toSlice())
}

// toSlice returns slice of set values.
func (s Set[T]) toSlice() []T {
	sl := make([]T, 0, len(s.values))
	for item := range s.values {
		sl = append(sl, item)
	}
	return sl
}

// Values returns slice of set values.
func (s *Set[T]) Values() []T { return s.toSlice() }

// Clone returns copy of current set.
func (s *Set[T]) Clone() *Set[T] {
	m := make(map[T]struct{}, s.Size())
	for item := range s.values {
		m[item] = struct{}{}
	}
	return &Set[T]{
		values: m,
	}
}

// Union performs a "union" on the sets and returns a new set.
// A union is a set of all elements that appear in either set. In mathmatics
// it's denoted as:
// 	A ∪ B
// Example:
// 	{1 2 3} ∪ {3 4 5} = {1 2 3 4 5}
// This operation is commutative, meaning you will get the same result no matter
// the order of the operands. In other words:
// 	setA.Union(setB) == setB.Union(setA)
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	res := s.Clone()
	for item := range other.values {
		res.values[item] = struct{}{}
	}
	return res
}

// Intersect performs an "intersection" on the sets and returns a new set.
// An intersection is a set of all elements that appear in both sets. In
// mathmatics it's denoted as:
// 	A ∩ B
// Example:
// 	{1 2 3} ∩ {3 4 5} = {3}
// This operation is commutative, meaning you will get the same result no matter
// the order of the operands. In other words:
// 	setA.Intersect(setB) == setB.Intersect(setA)
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	// use smaller set for optimization
	if s.Size() > other.Size() {
		return other.intersect(s)
	}
	return s.intersect(other)
}

// intersect returns intersection of two sets.
// This is a helper function for:
//  Set[T].Intersect(other *Set[T]) *Set[T]
// performing iteration over smaller set for optimization.
func (s *Set[T]) intersect(other *Set[T]) *Set[T] {
	res := NewEmptySet[T]()
	for item := range s.values {
		if other.Has(item) {
			_ = res.AddIfNotExist(item)
		}
	}
	return res
}

// Difference performs a "set difference" on the sets and returns a new set.
// A set difference resembles a subtraction, where the result is a set of all
// elements that appears in the first set but not in the second. In mathmatics
// it's denoted as:
// 	A \ B
// Example:
// 	{1 2 3} \ {3 4 5} = {1 2}
// This operation is noncommutative, meaning you will get different results
// depending on the order of the operands. In other words:
// 	setA.Difference(setB) != setB.Difference(setA)
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	res := NewEmptySet[T]()
	for item := range s.values {
		if !other.Has(item) {
			_ = res.AddIfNotExist(item)
		}
	}
	return res
}
