package calculatorservice

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

//var InternalServerErr = errors.New(" Result out of range")

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.Methods("GET").Path("/add/").Handler(httptransport.NewServer(
		e.AddEndpoint,
		decodeAddRequest,
		encodeResponse,
		options...,
	))
	return r
}
func encodeAddRequest(ctx context.Context, req *http.Request, request interface{}) error {

	req.URL.Path = "/add/"
	return encodeRequest(ctx, req, request)
}
func decodeAddResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response addResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	return response, err
}
func decodeAddRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var r addRequest
	err := json.NewDecoder(req.Body).Decode(&r)

	if err != nil {
		return nil, err
	}
	return r, nil
}

type errorer interface {
	error() error
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset= utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset= utf-8")
	w.WriteHeader(errorCodeFrom(err))
	s := err.Error()
	byteSlice := []byte(s)
	w.Write(byteSlice)
}

func errorCodeFrom(err error) int {
	switch err {
	case InternalServerErr:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

type addRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type addResponse struct {
	Result int   `json:"result"`
	Err    error `json:"err,omitempty"`
}
