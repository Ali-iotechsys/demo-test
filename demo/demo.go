package demo

import "errors"

var ErrInvalidSum1 = errors.New("invalid sum")

func Add(x, y int) (int, error) {
	if x <= 0 || y <= 0 {
		return 0, ErrInvalidSum1
	}
	return x + y, nil
}
