package errs

import (
	"errors"
	"github.com/bektosh03/pkg/models/responses"
	"github.com/gin-gonic/gin"
	"net/http"

	"google.golang.org/grpc/codes"
)

// Default Aborts

var (
	DefaultInternalAbort = Abort{
		When: codes.Internal,
		Then: func(c *gin.Context, msg string) {
			c.JSON(
				http.StatusInternalServerError,
				responses.ErrorResponse{
					Error:   codes.Internal.String(),
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
		Then: func(c *gin.Context, msg string) {
			c.JSON(
				http.StatusInternalServerError,
				responses.ErrorResponse{
					Error:   codes.Internal.String(),
					Message: msg,
				},
			)
		},
	}
)

var (
	ErrInternal = errors.New("something went wrong")
)
