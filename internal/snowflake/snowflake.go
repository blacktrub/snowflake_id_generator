package snowflake

import (
	"snowflake_id_generator/internal/conf"
	"snowflake_id_generator/internal/seq"
	"time"
)

type Snowflake struct {
	seq  *seq.Seq
	conf *conf.Conf
}

func (s *Snowflake) Get() (int64, error) {
	now := time.Now().Unix()
	n := now << 5
	n |= int64(s.conf.Dc)
	n = n << 5
	n |= int64(s.conf.Node)
	n = n << 12
	sq, err := s.seq.Next(now)
	if err != nil {
		return 0, err
	}
	n |= int64(sq)
	return n, nil
}

func New() (Snowflake, error) {
	seq := seq.New()
	conf, err := conf.New()
	if err != nil {
		return Snowflake{}, err
	}

	return Snowflake{seq: &seq, conf: &conf}, nil
}
