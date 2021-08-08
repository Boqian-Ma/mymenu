package errors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestErrorResponse_Error(t *testing.T) {
	e := ErrorResponse{
		Message: "abc",
	}

	assert.Equal(t, "abc", e.Error())
}

func TestErrorResponse_StatusCode(t *testing.T) {
	e := ErrorResponse{
		Status: 400,
	}

	assert.Equal(t, 400, e.StatusCode())
}

func TestAbort(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/invoices", nil)

	Abort(c, BadRequest(""))
	assert.True(t, c.IsAborted())
	assert.Len(t, c.Errors, 1)
	assert.Equal(t, "Error #01: Bad Request\n", c.Errors.String())
}

func TestSend(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/invoices", nil)

	Send(c, BadRequest(""))
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"message\":\"Bad Request\"}", w.Body.String())
}

func TestInternalServerError(t *testing.T) {
	res := InternalServerError("test")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode())
	assert.Equal(t, "test", res.Error())

	res = InternalServerError("")
	assert.NotEmpty(t, res.Error())
}

func TestNotFound(t *testing.T) {
	res := NotFound("test")
	assert.Equal(t, http.StatusNotFound, res.StatusCode())
	assert.Equal(t, "test", res.Error())

	res = NotFound("")
	assert.NotEmpty(t, res.Error())
}

func TestUnauthorized(t *testing.T) {
	res := Unauthorized("test")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode())
	assert.Equal(t, "test", res.Error())

	res = Unauthorized("")
	assert.NotEmpty(t, res.Error())
}

func TestForbidden(t *testing.T) {
	res := Forbidden("test")
	assert.Equal(t, http.StatusForbidden, res.StatusCode())
	assert.Equal(t, "test", res.Error())

	res = Forbidden("")
	assert.NotEmpty(t, res.Error())
}

func TestBadRequest(t *testing.T) {
	res := BadRequest("test")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	assert.Equal(t, "test", res.Error())

	res = BadRequest("")
	assert.NotEmpty(t, res.Error())
}

func TestNotImplemented(t *testing.T) {
	res := NotImplemented("test")
	assert.Equal(t, http.StatusNotImplemented, res.StatusCode())
	assert.Equal(t, "test", res.Error())

	res = NotImplemented("")
	assert.NotEmpty(t, res.Error())
}
