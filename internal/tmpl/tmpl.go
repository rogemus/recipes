package tmpl

import (
	"recipes.krogowski.dev/internal/ctx"
	"recipes.krogowski.dev/internal/session"
)

type Tmpl struct {
	Cache TmplCache
	*session.Session
	*ctx.Context
}

func New(session *session.Session, ctx *ctx.Context) Tmpl {
	return Tmpl{
		Session: session,
		Context: ctx,
	}
}
