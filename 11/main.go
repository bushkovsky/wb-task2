package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
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

func (s *Set[T]) printSet() {
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
	s1 := Set[int]{container: make(map[int]struct{})}
	s2 := Set[int]{container: make(map[int]struct{})}

	s1.container[1] = struct{}{}
	s1.container[3] = struct{}{}
	s2.container[2] = struct{}{}
	s2.container[1] = struct{}{}

	s3 := s1.Intersection(s2)
	s3.printSet()
}
