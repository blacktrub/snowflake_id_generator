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
	seq seq.Seq
}

func (s *Snowflake) Get() int64 {
	n := time.Now().Unix()
	n = n << 5
	n |= dc
	n = n << 5
	n |= node
	n = n << 12
	n |= int64(s.seq.Next())
	return n
}

func Init(seq seq.Seq) Snowflake {
	return Snowflake{seq: seq}
}
