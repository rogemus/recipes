package core

import (
	"log/slog"

	"recipies.krogowski.dev/internal/ctx"
	"recipies.krogowski.dev/internal/session"
	"recipies.krogowski.dev/internal/tmpl"
)

type Env struct {
	Ctx       *ctx.Context
	dupa      string
	Session   *session.Session
	Logger    *slog.Logger
	DebugMode bool
	TmplCache tmpl.TmplCache
}
