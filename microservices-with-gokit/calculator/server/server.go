package main

import (
	"github.com/amit1563/microservices-with-go/microservices-with-gokit/calculator"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestamp)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var s calculatorservice.Service
	{
		s = calculatorservice.NewService()
		s = calculatorservice.LoggingMiddleware(logger)(s)

	}

	addHandler := calculatorservice.MakeHTTPHandler(s, logger)

	http.ListenAndServe(":8080", addHandler)
}
