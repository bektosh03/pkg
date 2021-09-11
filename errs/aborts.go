package errs

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

type AbortionFunc func(c *gin.Context, err error, msg string)

// Abort defines an object for abortion in accordance to code.Code
// that is contained in status.Status (see https://google.golang.org/grpc/status#Status)
// returned from gRPC call.
type Abort struct {
	When codes.Code
	Then AbortionFunc
}

// Error defines an object for abortion in accordance to error which is matched to When.
// Should be used with any function that returns error (except from status.Status)
// which should be handled with Abort
type Error struct {
	When error
	Msg  string
	Then AbortionFunc
}

type Handler struct {
	aborts []Abort
	errs   []Error
}

type HandlerBuilder interface {
	WithAborts(aborts ...Abort) HandlerBuilder
	WithErrors(errs ...Error) HandlerBuilder
	Build() *Handler
}

func NewHandlerBuilder() HandlerBuilder {
	return new(Handler)
}

func (h *Handler) WithAborts(aborts ...Abort) HandlerBuilder {
	h.aborts = append(h.aborts, aborts...)
	return h
}

func (h *Handler) WithErrors(errs ...Error) HandlerBuilder {
	h.errs = append(h.errs, errs...)
	return h
}

func (h *Handler) Build() *Handler {
	return h
}
