package httpserv

import "github.com/Skillbox_30_2023_new/internal/usecase"

type HTTPHandler struct {
	Service *usecase.UserService
}

func NewHTTPHandler(service *usecase.UserService) *HTTPHandler {
	return &HTTPHandler{
		Service: service,
	}
}
