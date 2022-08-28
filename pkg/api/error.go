package api

import "fmt"

type ErrorCode int

var (
	// ErrorCodeInternalServerError is returned when there's an internal error that must be retried
	ErrorCodeInternalServerError ErrorCode = 10500

	// ErrorCodeInvalidContentType when the request content type is invalid, not JSON for the moment
	ErrorCodeInvalidContentType ErrorCode = 10415

	// ErrorCodeBadRequest when the request is invalid
	ErrorCodeBadRequest ErrorCode = 10400

	// ErrorCodeResourceNotFound when the resource specified was not found
	ErrorCodeResourceNotFound ErrorCode = 10404
)

// ApiError struct returned to the user when bad request or error
type ErrorResponse struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

var (
	// InternalServerError when api failed processing the request internally
	InternalServerError = &ErrorResponse{
		Code:    ErrorCodeInternalServerError,
		Message: "internal server error, try again later",
	}

	// ResourceNotFoundError when api failed to found the specified resource
	ResourceNotFoundError = &ErrorResponse{
		Code:    ErrorCodeResourceNotFound,
		Message: "specified resource not found",
	}
)

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    ErrorCodeBadRequest,
		Message: "invalid request: " + err.Error(),
	}
}

// Error
func (apiErr ErrorResponse) Error() string {
	return fmt.Sprintf("%s (%d)", apiErr.Message, apiErr.Code)
}
