package demo

import "errors"

var ErrInvalidSum = errors.New("invalid sum")

func Add(x, y int) (int, error) {
	if x <= 0 || y <= 0 {
		return 0, ErrInvalidSum
	}
	return x + y, nil
}
