package services

import "sync"

// NumberGen — генератор инвентарных номеров.
type NumberGen interface {
	Next() int
}

type seq struct {
	mu sync.Mutex
	n  int
}

func NewSeq() NumberGen { return &seq{} }

func (s *seq) Next() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.n++
	return s.n
}
