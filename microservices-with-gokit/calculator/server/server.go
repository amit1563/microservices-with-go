package main

import (
	"flag"
	"fmt"
	"github.com/amit1563/microservices-with-go/microservices-with-gokit/calculator"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

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
	// Interrupt handler.
	err := make(chan error)
	go func() {
		osSigChannel := make(chan os.Signal)
		signal.Notify(osSigChannel, syscall.SIGINT, syscall.SIGTERM)
		err <- fmt.Errorf("%s", <-osSigChannel)
	}()
	// HTTP transport.
	go func() {
		err <- http.ListenAndServe(*httpAddr, httpHandler)
	}()
	// Run!
	logger.Log("lastcall", <-err)

}
