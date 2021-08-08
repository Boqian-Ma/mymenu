package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPError extends the error interface to allow setting a status code
type HTTPError interface {
	Error() string
	StatusCode() int
}

// ErrorResponse is the response that represents an error
type ErrorResponse struct {
	Status  int    `json:"-"`
	Message string `json:"message" example:"error message"`
}

// Error is required by the error interface
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is used to create a HTTPError
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

// Abort attaches the given error to the ctx then aborts the request
func Abort(ctx *gin.Context, e error) {
	ctx.Error(e)
	ctx.Abort()
}

// Send converts a HTTPErorr to a HTTP resp and attaches it to the request context
func Send(ctx *gin.Context, e HTTPError) {
	ctx.JSON(e.StatusCode(), gin.H{"message": e.Error()})
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "Bad Request"
	}

	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

// Forbidden creates a new error response representing a forbidden request (HTTP 403)
func Forbidden(msg string) ErrorResponse {
	if msg == "" {
		msg = "Forbidden"
	}

	return ErrorResponse{
		Status:  http.StatusForbidden,
		Message: msg,
	}
}

// Unauthorized creates a new error response representing an authorized request (HTTP 401)
func Unauthorized(msg string) ErrorResponse {
	if msg == "" {
		msg = "Unauthorized"
	}

	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: msg,
	}
}

// NotFound creates a new error response representing a not found request (HTTP 404)
func NotFound(msg string) ErrorResponse {
	if msg == "" {
		msg = "Not Found"
	}

	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

// Conflict creates a new error response representing a conflict request (HTTP 409)
func Conflict(msg string) ErrorResponse {
	if msg == "" {
		msg = "Conflict"
	}

	return ErrorResponse{
		Status:  http.StatusConflict,
		Message: msg,
	}
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) ErrorResponse {
	if msg == "" {
		msg = "Something went wrong, please try again later"
	}

	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

// NotImplemented creates a new error response representing a not implemented error (HTTP 502)
func NotImplemented(msg string) ErrorResponse {
	if msg == "" {
		msg = "Not implemented"
	}

	return ErrorResponse{
		Status:  http.StatusNotImplemented,
		Message: msg,
	}
}
