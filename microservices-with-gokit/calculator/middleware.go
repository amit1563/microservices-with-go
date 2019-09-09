package calculatorservice

import (
	"github.com/go-kit/kit/log"
	"time"
)

type Middleware func(s Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) Add(x int, y int) (result int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("Method", "Add", "TimeTook", "Result", result, time.Since(begin), "error", err)
	}(time.Now())
	return mw.next.Add(x, y)
}
