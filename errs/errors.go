package errs

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Check(err error, c *gin.Context) (yes bool) {
	if err == nil {
		return false
	}
	er := matchError(h.errs, err)
	er.Then(c, er.Msg)
	return true
}

func matchError(errs []Error, err error) Error {
	for _, er := range errs {
		if errors.Is(err, er.When) {
			return er
		}
	}
	return DefaultInternalError
}
