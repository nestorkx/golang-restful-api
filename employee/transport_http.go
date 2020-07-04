package employee

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"golang-restful-api/helper"
	"net/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getEmployeesHandler := kithttp.NewServer(makeGetEmployeesEndpoint(s), getEmployeesRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	return r
}

func getEmployeesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
