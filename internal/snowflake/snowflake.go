package snowflake

import (
	"snowflake_id_generator/internal/seq"
	"time"
)

// TODO: get those numbers from ENV
// TODO: must be from 0 to 31
const dc = 31
const node = 31

type Snowflake struct {
	seq *seq.Seq
}

func (s *Snowflake) Get() (int64, error) {
	now := time.Now().Unix()
	n := now << 5
	n |= dc
	n = n << 5
	n |= node
	n = n << 12
	sq, err := s.seq.Next(now)
	if err != nil {
		return 0, err
	}
	n |= int64(sq)
	return n, nil
}

func New() Snowflake {
	seq := seq.New()
	return Snowflake{seq: &seq}
}
