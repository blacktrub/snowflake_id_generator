package conf

// TODO: get those numbers from ENV
// TODO: must be from 0 to 31
const dc = 31
const node = 31

type Conf struct {
	Dc   int
	Node int
}

func New() Conf {
	return Conf{Dc: dc, Node: node}
}
