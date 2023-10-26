package service

import (
	"sort"
)

// IntOrderedSet structure
type IntOrderedSet struct {
	data map[int]struct{}
}

// NewIntOrderedSet creates a new IntOrderedSet
func NewIntOrderedSet() *IntOrderedSet {
	return &IntOrderedSet{
		data: make(map[int]struct{}),
	}
}

// Add adds an integer to the set and ensures the set remains sorted
func (s *IntOrderedSet) Add(item int) {
	s.data[item] = struct{}{}
}

// Remove removes an integer from the set
func (s *IntOrderedSet) Remove(item int) {
	delete(s.data, item)
}

// Contains checks if an integer exists in the set
func (s *IntOrderedSet) Contains(item int) bool {
	_, exists := s.data[item]
	return exists
}

// Items returns all integers in the set in sorted order
func (s *IntOrderedSet) Items() []int {
	keys := make([]int, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

// Size returns the number of items in the set
func (s *IntOrderedSet) Size() int {
	return len(s.data)
}
