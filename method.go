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
	handler http.Handler
}

func (h MethodHandler) Options() []string {
	return h.methods
}

func (h MethodHandler) Error() string {
	return fmt.Sprintf("Allow: %s", strings.Join(h.methods, ", "))
}

func (h MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		h.handler.ServeHTTP(w, r)
		return
	}

	for _, method := range h.methods {
		if method == r.Method {
			h.handler.ServeHTTP(w, r)
			return
		}
	}

	panic(h) // 405
}

// Adds path only supports GET to Router
func (r *Router) GET(path string, handler http.HandlerFunc) {
	h := MethodHandler{
		methods: []string{"GET", "HEAD"},
		handler: handler,
	}
	r.Handle(path, h)
}

// MethodHandlerFunc makes a Handler only accept a given set of Methods
func MethodHandlerFunc(methods []string, handler http.HandlerFunc) http.Handler {
	return MethodHandler{
		methods: methods,
		handler: handler,
	}
}
