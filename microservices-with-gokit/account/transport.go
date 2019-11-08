package account

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	logger.Log("called", "MakeHTTPHandler")
	r := mux.NewRouter()
	e := MakeServerEnpoint(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	r.Methods("POST").Path("/signUp").Handler(httptransport.NewServer(
		e.SignUpEndpoint,
		decodeSignUpRequest,
		encodeResponse,
		options...,
	))
	return r
}
func decodeSignUpRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var signUpRequest SignUpRequest
	json.NewDecoder(req.Body).Decode(&signUpRequest)
	return signUpRequest, nil
}
func encodeResponse(ctx context.Context, w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset= utf-8")
	return json.NewEncoder(w).Encode(res)
}

type SignUpRequest struct {
	SignUp SignUp `json:"signUp"`
}
type SignUpResponse struct {
	SignUp string `json:"msg"`
	Err    error  `json:"error,omitempty"`
}
