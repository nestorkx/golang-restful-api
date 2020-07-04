package employee

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"golang-restful-api/helper"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}

func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {
	getEmployeesEnpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		res, err := s.GetEmployees(&req)
		helper.Catch(err)
		return res, nil
	}
	return getEmployeesEnpoint
}
