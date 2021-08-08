package errors

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// we don't care about broken pipp errors
func isBrokenPipeError(err interface{}) bool {
	if netErr, ok := err.(*net.OpError); ok {
		if sysErr, ok := netErr.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(sysErr.Error()), "broken pipe") ||
				strings.Contains(strings.ToLower(sysErr.Error()), "connection reset by peer") {
				return true
			}
		}
	}

	return false
}

func Handler(serviceName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if recoverErr := recover(); recoverErr != nil {
				Send(ctx, InternalServerError(""))

				err, ok := recoverErr.(error)
				if !ok {
					err = errors.New("unknown panic")
				}

				if !isBrokenPipeError(recoverErr) {
					fmt.Println("error: caught panic " + err.Error())
				} else {
					fmt.Println("debug: ignoring broken pipe error")
				}
			}
		}()

		ctx.Next()

		err := ctx.Errors.Last()
		if err == nil {
			return
		}

		if errorResponse, ok := err.Err.(HTTPError); ok && errorResponse.StatusCode() < 500 {
			fmt.Println("info: sending client error " + errorResponse.Error())
			Send(ctx, errorResponse)
		} else {
			fmt.Println("error: internal server error " + errorResponse.Error())
			Send(ctx, InternalServerError(""))
		}
	}
}
