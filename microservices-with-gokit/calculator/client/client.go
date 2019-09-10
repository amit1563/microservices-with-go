package client

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	consulapi "github.com/hashicorp/consul/api"
	"io"
	"time"

	"github.com/amit1563/microservices-with-go/microservices-with-gokit/calculator"
	"github.com/go-kit/kit/log"
)

func New(serverAddr string, logger log.Logger) (calculatorservice.Service, error) {
	apiClient, err := consulapi.NewClient(&consulapi.Config{
		Address: serverAddr,
	})
	if err != nil {
		return nil, err
	}
	var (
		consulService = "calculatorservice"
		consulTags    = []string{"DEV"}
	)
	var (
		sdclient  = consul.NewClient(apiClient)
		instancer = consul.NewInstancer(sdclient, logger, consulService, consulTags, true)
		endpoints calculatorservice.Endpoints
	)
	{
		factory := factoryFor(calculatorservice.MakeAddEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddEndpoint = retry
	}
	return endpoints, nil
}

func factoryFor(makeEndpoint func(calculatorservice.Service) endpoint.Endpoint) sd.Factory {
	return func(serverAddr string) (endpoint.Endpoint, io.Closer, error) {
		service, err := calculatorservice.MakeClientEndpoints(serverAddr)
		if err != nil {
			return nil, nil, err
		}
		return makeEndpoint(service), nil, nil
	}
}
