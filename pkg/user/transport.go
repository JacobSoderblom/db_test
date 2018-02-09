package user

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"context"
	"encoding/json"
	"errors"
	"strconv"
)

func MakeHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(httptransport.DefaultErrorEncoder),
	}

	{
		handler := httptransport.NewServer(
			makeCreateEndpoint(s),
			decodeCreateRequest,
			httptransport.EncodeJSONResponse,
			options...,
		)

		r.Methods("POST").Path("/api/user/").Handler(handler)
	}

	{
		handler := httptransport.NewServer(
			makeListEndpoint(s),
			decodeListRequest,
			httptransport.EncodeJSONResponse,
			options...,
		)

		r.Methods("GET").Path("/api/user/all").Handler(handler)
	}

	{
		handler := httptransport.NewServer(
			makeFindEndpoint(s),
			decodeFindRequest,
			httptransport.EncodeJSONResponse,
			options...,
		)

		r.Methods("GET").Path("/api/user/{id}").Handler(handler)
	}

	return r
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Street   string `json:"street"`
		City     string `json:"city"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return createRequest{
		body.Name,
		body.Email,
		body.Password,
		body.Street,
		body.City,
	}, nil
}

func decodeFindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("unknown user")
	}

	uID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("unknown user")
	}

	return uID, nil
}

func decodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
