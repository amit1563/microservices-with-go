package calculatorservice

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		AddEndpoint: MakeAddEndpoint(s),
	}
}
func MakeAddEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		//res := (addResponse)
		r, err := s.Add(req.X, req.Y)
		if err != nil {
			return addResponse{0, err}, err
		}
		return addResponse{r, nil}, nil
	}
}
