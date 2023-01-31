package conf

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Conf struct {
	Dc   int
	Node int
}

func checkValue(x int) bool {
	if x >= 0 && x < 32 {
		return true
	}
	return false
}

func parseEnv(name string) (int, error) {
	rval, exists := os.LookupEnv(name)
	if !exists {
		return 0, errors.New(fmt.Sprintf("%v is not set", name))
	}

	val, err := strconv.Atoi(rval)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("%v must be integer", name))
	}

	if !checkValue(val) {
		return 0, errors.New(fmt.Sprintf("%v must be in range from 0 to 31", name))
	}
	return val, nil
}

func New() (Conf, error) {
	dc, err := parseEnv("SNOWFLAKE_DC")
	if err != nil {
		return Conf{}, err
	}

	node, err := parseEnv("SNOWFLAKE_NODE")
	if err != nil {
		return Conf{}, err
	}

	return Conf{Dc: dc, Node: node}, nil
}
