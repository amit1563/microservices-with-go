package main

import (
	"flag"
	"github.com/amit1563/microservices-with-go/microservices-with-gokit/account"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
)

func main() {
	var httpAddr = flag.String("httpAddr", ":8080", "http Listen address host:port")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"caller", log.DefaultCaller,
			"service", "account-service",
			//"timestamp", log.DefaultTimestampUTC,
		)
	}
	var repo account.Repoisitory

	// create bolt db

	db, _ := account.SetupBolt()
	repo = account.NewRepo(db, logger)

	var s account.Service
	s = account.NewService(repo, logger)

	var httpHandler http.Handler
	{
		httpHandler = account.MakeHTTPHandler(s, log.With(logger, "component", "http"))
	}
	flag.Parse()
	logger.Log("Starting", "Server")
	defer logger.Log("Server stopped")
	http.ListenAndServe(*httpAddr, httpHandler)

}
