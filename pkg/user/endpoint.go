package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type createRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Street   string `json:"street"`
	City     string `json:"city"`
}

type userResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Street string `json:"street"`
	City   string `json:"city"`
}

type listResponse struct {
	Data []*userListItem `json:"data"`
}

func makeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)

		u, err := s.Create(req.Name, req.Email, req.Password, Address{req.Street, req.City})
		if err != nil {
			return nil, err
		}

		return userResponse{
			int(u.ID),
			u.Name,
			u.Email,
			u.Address.Street,
			u.Address.City,
		}, nil
	}
}

func makeFindEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		id := request.(int)

		u, err := s.Find(UserID(id))
		if err != nil {
			return nil, err
		}

		return userResponse{
			int(u.ID),
			u.Name,
			u.Email,
			u.Address.Street,
			u.Address.City,
		}, nil
	}
}

func makeListEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		list, err := s.All()
		if err != nil {
			return nil, err
		}

		return listResponse{
			Data: list,
		}, nil
	}
}
