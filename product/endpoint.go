package product

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

func makeGetProductByIdEndpoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}

	return getProductByIdEndpoint
}

func makeGetProductsEndpoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		res, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return res, nil
	}
	return getProductsEndPoint
}

func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productID, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}
		return productID, nil
	}
	return addProductEndpoint
}
