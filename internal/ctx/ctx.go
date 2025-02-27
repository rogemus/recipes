package ctx

import (
	"context"
	"net/http"
)

type contextKey string

const IsAuthenticatedContextKey = contextKey("isAuthenticated")
const UserIdContextKey = contextKey("userId")

type Context struct{}

func New() *Context {
	return &Context{}
}

func (c *Context) IsAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(IsAuthenticatedContextKey).(bool)

	if !ok {
		return false
	}

	return isAuthenticated
}

func (c *Context) CopyCtxWithKey(r *http.Request, key contextKey, value any) context.Context {
	ctx := context.WithValue(r.Context(), key, value)
	return ctx
}
