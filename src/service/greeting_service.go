package main

import (
	"context"
	"strings"
	"encoding/json"
	"net/http"
	"github.com/go-kit/kit/endpoint"

	transport "github.com/go-kit/kit/transport/http"
)

type GreetingService interface{
	greet(ctx context.Context, s string) (string, error)
}

type greetingService struct {

}

type greetNameRequest struct {
	S string `json:"s"`
}

func (greetingService) greet(_ context.Context, s string) (string, error) {
		return strings.ToUpper("Hello " + s), nil
}


func greetingEndPoint(svc GreetingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(greetNameRequest)
		v, err := svc.greet(ctx,req.S)
		if err != nil {
			return err, nil
		}
		return v,nil
	}
}


func main() {
	svc := greetingService{}

	greetHandler := transport.NewServer(greetingEndPoint(svc),decodeRequest,encodeResponse)

	http.Handle("/greet", greetHandler)

	http.ListenAndServe(":8081", nil)
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request greetNameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
