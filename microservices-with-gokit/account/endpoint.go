package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	SignUpEndpoint endpoint.Endpoint
}

func MakeServerEnpoint(s Service) Endpoint {
	return Endpoint{
		SignUpEndpoint: MakeSignUpEndpoint(s),
	}
}
func MakeSignUpEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(SignUpRequest)

		user := User{
			Name: "Test",
			Account: Account{
				Username: request.SignUp.Username,
				Password: request.SignUp.Password,
				Email:    request.SignUp.Email,
			},
		}

		err := s.Create(ctx, user)
		if err != nil {
			return nil, err
		}
		return SignUpResponse{SignUp: "successfully signed up"}, nil
	}
}
