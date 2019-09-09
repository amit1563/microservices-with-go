package main

import (
	"github.com/amit1563/microservices-with-go/microservices-with-gokit/calculator"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
)

func main() {

	var s calculatorservice.Service
	s = calculatorservice.NewService()

	addHandler := calculatorservice.MakeHTTPHandler(s, log.NewLogfmtLogger(os.Stderr))

	http.ListenAndServe(":8080", addHandler)
}
