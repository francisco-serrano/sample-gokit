package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

type HelloService interface {
	Hello(v string) string
}

type helloService struct {
}

func (s helloService) Hello(v string) string {
	return "hello " + v
}

type helloRequest struct {
	V string `json:"v"`
}

type helloResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func makeHelloEndpoint(svc HelloService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(helloRequest)
		if !ok {
			return helloResponse{
				Error: errors.New("invalid request structure").Error(),
			}, nil
		}

		return helloResponse{
			Message: svc.Hello(req.V),
		}, nil
	}
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request helloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := helloService{}

	helloHandler := httptransport.NewServer(
		makeHelloEndpoint(svc),
		decodeHelloRequest,
		encodeResponse,
	)

	http.Handle("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
