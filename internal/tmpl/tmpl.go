package tmpl

import (
	"recipies.krogowski.dev/internal/ctx"
	"recipies.krogowski.dev/internal/session"
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
