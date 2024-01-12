package httpserv

import "github.com/Skillbox_30_2023_new/internal/usecase"

// HTTPHandler is a HTTP handler for the UserService.
type HTTPHandler struct {
	Service *usecase.UserService
}

// NewHTTPHandler creates a new HTTPHandler.
func NewHTTPHandler(service *usecase.UserService) *HTTPHandler {
	return &HTTPHandler{
		Service: service,
	}
}
