package service

import (
	"sort"
)

type FloorRequest struct {
	Source      int
	Destination int
}

// IntOrderedSet structure
type OrderedSet struct {
	data  map[FloorRequest]struct{}
	items []FloorRequest
}

// NewIntOrderedSet creates a new IntOrderedSet
func NewOrderedSet() *OrderedSet {
	return &OrderedSet{
		data:  make(map[FloorRequest]struct{}),
		items: []FloorRequest{},
	}
}

// Add adds an integer to the set and ensures the set remains sorted
func (s *OrderedSet) Add(req FloorRequest) {
	if _, exists := s.data[req]; !exists {
		s.data[req] = struct{}{}
		index := sort.Search(len(s.items), func(i int) bool { return s.items[i].Source >= req.Source })
		s.items = append(s.items[:index], append([]FloorRequest{req}, s.items[index:]...)...)
	}
}

// Remove removes an integer from the set
func (s *OrderedSet) Remove(req FloorRequest) {
	if _, exists := s.data[req]; exists {
		delete(s.data, req)
		index := sort.Search(len(s.items), func(i int) bool { return s.items[i].Source == req.Source })
		if index < len(s.items) && s.items[index].Source == req.Source {
			s.items = append(s.items[:index], s.items[index+1:]...)
		}
	}
}

// Contains checks if an integer exists in the set
func (s *OrderedSet) Contains(item FloorRequest) bool {
	_, exists := s.data[item]
	return exists
}

// Items returns all integers in the set in sorted order

func (s *OrderedSet) Items() []FloorRequest {
	return s.items
}

// Size returns the number of items in the set
func (s *OrderedSet) Size() int {
	return len(s.data)
}
