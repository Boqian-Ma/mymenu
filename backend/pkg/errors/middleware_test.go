package errors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	t.Run("normal processing", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, router := gin.CreateTestContext(w)
		router.Use(Handler(""))
		router.GET("/errors_test", handlerOK)
		c.Request, _ = http.NewRequest("GET", "/errors_test", nil)

		router.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "{\"message\":\"success\"}", w.Body.String())
		assert.False(t, c.IsAborted())
		assert.Len(t, c.Errors, 0)
	})

	t.Run("HTTP error processing", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, router := gin.CreateTestContext(w)
		router.Use(Handler(""))
		router.GET("/errors_test", handlerHTTPError)
		c.Request, _ = http.NewRequest("GET", "/errors_test", nil)

		router.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"message\":\""+BadRequest("").Message+"\"}", w.Body.String())
	})

	t.Run("error processing", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, router := gin.CreateTestContext(w)
		router.Use(Handler(""))
		router.GET("/errors_test", handlerError)
		c.Request, _ = http.NewRequest("GET", "/errors_test", nil)

		router.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "{\"message\":\""+InternalServerError("").Message+"\"}", w.Body.String())
	})

	t.Run("panic processing", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, router := gin.CreateTestContext(w)
		router.Use(Handler(""))
		router.GET("/errors_test", handlerPanic)
		c.Request, _ = http.NewRequest("GET", "/errors_test", nil)

		router.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "{\"message\":\""+InternalServerError("").Message+"\"}", w.Body.String())
	})
}

func handlerOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func handlerError(c *gin.Context) {
	// This is how errors are handled by the API layer
	Abort(c, fmt.Errorf("abc"))
}

func handlerHTTPError(c *gin.Context) {
	Abort(c, BadRequest(""))
}

func handlerPanic(c *gin.Context) {
	panic("xyz")
}
