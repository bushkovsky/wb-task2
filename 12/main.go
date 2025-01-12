package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

type Set[T comparable] struct {
	container map[T]struct{}
}

func (s *Set[T]) Equal(a, b T) bool {
	return a == b
}

func (s *Set[T]) Union(second Set[T]) {
	for i := range second.container {
		s.container[i] = struct{}{}
	}
}

func (s *Set[T]) Add(value T) {
	s.container[value] = struct{}{}
}

func (s *Set[T]) PrintSet() {
	for i := range s.container {
		fmt.Println(i)
	}
}

func (s *Set[T]) Intersection(second Set[T]) (intersection Set[T]) {
	intersection.container = make(map[T]struct{})

	for i := range second.container {
		if _, exists := s.container[i]; exists {
			intersection.container[i] = struct{}{}
		}
	}
	return intersection
}

func main() {
	// Последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание множества
	set := Set[string]{container: make(map[string]struct{})}

	// Добавление строк в множество
	for _, item := range sequence {
		set.Add(item)
	}

	// Печать множества
	set.PrintSet()
}
