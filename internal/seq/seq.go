package seq

import (
	"errors"
	"sync"
)

const maxSeqInt = 4096

type Seq struct {
	n   int
	m   *sync.Mutex
	cur int64
}

func (s *Seq) Next(t int64) (int, error) {
	s.m.Lock()
	defer s.m.Unlock()

	// TODO: we need UNIX milliseconds here
	if t != s.cur {
		s.n = 0
		s.cur = t
		return s.n, nil
	}

	s.n++

	if s.n >= maxSeqInt {
		return 0, errors.New("Reached max ID per millisecond")
	}
	return s.n, nil
}

func Init() Seq {
	mut := sync.Mutex{}
	return Seq{m: &mut}
}
