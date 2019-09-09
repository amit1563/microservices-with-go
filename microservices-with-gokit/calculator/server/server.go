package main

import (
	"fmt"
	"github.com/amit1563/microservices-with-go/microservices-with-gokit/calculator"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	var httpHandler http.Handler
	{
		httpHandler = calculatorservice.MakeHTTPHandler(s, log.With(logger, "component", "http"))
	}
	err := make(chan error)
	go func() {
		osSigChannel := make(chan os.Signal)
		signal.Notify(osSigChannel, syscall.SIGINT, syscall.SIGTERM)
		err <- fmt.Errorf("%s", <-osSigChannel)
	}()
	go func() {
		err <- http.ListenAndServe(":8080", httpHandler)
	}()
	logger.Log("lastcall", <-err)
}
