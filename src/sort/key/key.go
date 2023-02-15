package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score float32
}

type By func(s1, s2 *Student) bool

type studentSorter struct {
	data []Student
	by   func(s1, s2 *Student) bool
}

func main() {
	s := []Student{
		{"Maria", 39.3},
		{"Doyle", 94.3},
		{"Cary", 11.3},
	}

	name := func(p1, p2 *Student) bool {
		return p1.name < p2.name
	}

	score := func(p1, p2 *Student) bool {
		return p1.score < p2.score
	}

	reverseScore := func(p1, p2 *Student) bool {
		return !score(p1, p2)
	}

	By(name).Sort(s)
	fmt.Println(s)

	By(score).Sort(s)
	fmt.Println(s)

	By(reverseScore).Sort(s)
	fmt.Println(s)
}

func (by By) Sort(students []Student) {
	sorter := &studentSorter{
		students,
		by,
	}

	sort.Sort(sorter)
}

func (s *studentSorter) Len() int {
	return len(s.data)
}

func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.data[i], &s.data[j])
}

func (s *studentSorter) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}
