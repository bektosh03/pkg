package errs

import (
	"errors"
	"github.com/bektosh03/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"

	"google.golang.org/grpc/codes"
)

// Default errors

var (
	ErrInternal = errors.New("something went wrong")
	ErrBadPageValue = errors.New("bad value for page")
	ErrBadLimitValue = errors.New("bad value for limit")

	ErrBadEmail = errors.New("bad value for email")
)

// Default Aborts

var (
	DefaultInternalAbort = Abort{
		When: codes.Internal,
		Then: func(c *gin.Context, err error, msg string) {
			c.JSON(
				http.StatusInternalServerError,
				models.ErrorResponse{
					Error:   err.Error(),
					Message: msg,
				},
			)
		},
	}
)

// Default Errors

var (
	DefaultInternalError = Error{
		When: ErrInternal,
		Msg: "something went wrong",
		Then: func(c *gin.Context, err error, msg string) {
			c.JSON(
				http.StatusInternalServerError,
				models.ErrorResponse{
					Error:   err.Error(),
					Message: msg,
				},
			)
		},
	}
	DefaultBadPageValueError = Error{
		When: ErrBadPageValue,
		Msg: "bad value for page, must be a positive integer",
		Then: func(c *gin.Context, err error, msg string) {
			c.JSON(
				http.StatusBadRequest,
				models.ErrorResponse{
					Error:   err.Error(),
					Message: msg,
				},
			)
		},
	}
	DefaultBadLimitValueError = Error{
		When: ErrBadPageValue,
		Msg: "bad value for page, must be a positive integer",
		Then: func(c *gin.Context, err error, msg string) {
			c.JSON(
				http.StatusBadRequest,
				models.ErrorResponse{
					Error:   err.Error(),
					Message: msg,
				},
			)
		},
	}
	DefaultBadEmailError = Error{
		When: ErrBadEmail,
		Msg: "invalid email, must be of form : 'user@example.com'",
		Then: func(c *gin.Context, err error, msg string) {
			c.JSON(
				http.StatusBadRequest,
				models.ErrorResponse{
					Error:   err.Error(),
					Message: msg,
				},
			)
		},
	}
)
