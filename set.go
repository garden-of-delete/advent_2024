package main

import "fmt"

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.elements))
	for key := range s.elements {
		values = append(values, key)
	}
	return values
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *Set[T]) Intersects(t *Set[T]) bool {
	for _, v := range s.Values() {
		if t.Contains(v) {
			return true
		}
	}
	return false
}

func runSetTests() {

	set := NewSet[int]()
	set2 := NewSet[int]()

	// Adding elements
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(3)
	set2.Add(1)
	set2.Add(9)
	set2.Add(7)

	fmt.Println("Set contains:", set.Values())
	fmt.Println("Set2 contains:", set2.Values())

	// Checking membership
	fmt.Println("Contains 2:", set.Contains(2))
	fmt.Println("Contains 4:", set.Contains(4))

	// Removing an element
	set.Remove(2)
	fmt.Println("After removing 2:", set.Values())

	// Checking intersect
	fmt.Println("Set1 and Set2 intersect: ", set.Intersects(set2))
	set2.Remove(1)
	fmt.Println("After removing 1: ", set.Intersects(set2))

	// Checking size
	fmt.Println("Set size:", set.Size())

	// Checking if empty
	fmt.Println("Is empty:", set.IsEmpty())

	// Clearing the set
	set.Clear()
	fmt.Println("After clearing:", set.Values())
	fmt.Println("Is empty:", set.IsEmpty())
}
