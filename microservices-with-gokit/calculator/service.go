package calculatorservice

import (
	"context"
	"errors"
	"math"
)

var (
	InternalServerErr = errors.New("Bad Request")
)

type Service interface {
	Add(context.Context, int, int) (int, error)
}

type service struct {
}

func NewService() Service {
	return service{}
}

func (service) Add(_ context.Context, x int, y int) (int, error) {
	if x > math.MaxInt16 || x-y > math.MaxInt16 {
		return 0, InternalServerErr
	}
	return x + y, nil
}
