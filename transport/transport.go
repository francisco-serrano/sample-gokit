package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/francisco-serrano/sample-gokit/service"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"strings"
)

type helloRequest struct {
	V string `json:"v"`
}

type helloResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func MakeHelloEndpoint(svc service.HelloService) endpoint.Endpoint {
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

func DecodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request helloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	if strings.TrimSpace(request.V) == "" {
		return nil, errors.New("value cannot be empty")
	}

	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
