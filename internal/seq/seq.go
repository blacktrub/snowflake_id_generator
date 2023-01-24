package seq

import "sync"

const maxSeqInt = 4096

type Seq struct {
	n int
	m *sync.Mutex
}

func (s *Seq) Next() int {
	s.m.Lock()
	defer s.m.Unlock()
	s.n++

	// TODO: we need to start from zero each millisecond
	if s.n >= maxSeqInt {
		s.n = 0
	}

	return s.n
}

func Init() Seq {
	mut := sync.Mutex{}
	return Seq{m: &mut}
}
