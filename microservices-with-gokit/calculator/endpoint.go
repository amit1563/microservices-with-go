package calculatorservice

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/url"
	"strings"
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
		r, err := s.Add(ctx, req.X, req.Y)
		if err != nil {
			return addResponse{0, err}, err
		}
		return addResponse{r, nil}, nil
	}
}

func MakeClientEndpoints(target string) (Endpoints, error) {
	if !strings.HasPrefix(target, "http") {
		target = "http://" + target
	}

	tgt, err := url.Parse(target)

	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	return Endpoints{
		AddEndpoint: httptransport.NewClient("GET", tgt, encodeAddRequest, decodeAddResponse, options...).Endpoint(),
	}, nil
}

// implement client endpoint

func (e Endpoints) Add(ctx context.Context, x int, y int) (int, error) {
	req := addRequest{X: x, Y: y}
	res, err := e.AddEndpoint(ctx, req)
	if err != nil {
		return 0, nil
	}
	response := res.(addResponse)
	return response.Result, nil
}
