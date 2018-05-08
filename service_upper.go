package main

import (
	"context"
	"strings"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"encoding/json"

	httptransport "github.com/go-kit/kit/transport/http"
	"log"
)

type StringService interface {
	Uppercase(context.Context, string) (string, error)

}

var ErrEmpty = errors.New("empty string")

type stringService struct{}

func (stringService) Uppercase(_ context.Context, s string) (string, error) {
	if  "" == strings.TrimSpace(s){
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}



func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}


func main(){
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}



func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
