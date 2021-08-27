package errs

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CheckGrpc(err error, c *gin.Context) (aborted bool) {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	if ok {
		matchStatusCode(h.aborts, st.Code()).Then(c, st.Message())
		return true
	}
	DefaultInternalAbort.Then(c, "something went wrong")
	return true
}

func matchStatusCode(aborts []Abort, code codes.Code) Abort {
	for _, abort := range aborts {
		if code == abort.When {
			return abort
		}
	}
	return DefaultInternalAbort
}
