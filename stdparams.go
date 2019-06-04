package mux

import (
	"context"
)

type paramsKey struct{}

// ParamsKey is the request context key under which URL params are stored.
//
var ParamsKey = paramsKey{}

// ParamsFromContext pulls the URL parameters from a request context,
// or returns nil if none are present.
func ParamsFromContext(ctx context.Context) Params {
	p, _ := ctx.Value(ParamsKey).(Params)
	return p
}
