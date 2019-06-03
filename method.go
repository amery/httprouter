package mux

import (
	"fmt"
	"net/http"
	"strings"
)

// MethodNotAllowed is an error that also provides a list of the
// supported Methods
type MethodNotAllowed interface {
	Options() []string
	Error() string
}

// MethodHandler is an http.Handler and a MethodNotAllowed error
// calling another handler only if the Method is right
type MethodHandler struct {
	methods []string
	handle  Handle
}

func (h MethodHandler) Options() []string {
	return h.methods
}

func (h MethodHandler) Error() string {
	return fmt.Sprintf("Allow: %s", strings.Join(h.methods, ", "))
}

func (h MethodHandler) Handle(w http.ResponseWriter, r *http.Request, ps Params) {
	if r == nil {
		h.handle(w, r, ps)
		return
	}

	for _, method := range h.methods {
		if method == r.Method {
			h.handle(w, r, ps)
			return
		}
	}

	panic(h) // 405
}

// Adds path only supports GET to Router
func (r *Router) GET(path string, handle Handle) {
	h := MethodHandler{
		methods: []string{"GET", "HEAD"},
		handle:  handle,
	}
	r.Handle(path, h.Handle)
}
