package core

import (
	"log/slog"

	"recipies.krogowski.dev/internal/ctx"
	"recipies.krogowski.dev/internal/session"
	"recipies.krogowski.dev/internal/tmpl"
)

type Env struct {
	Logger    *slog.Logger
	DebugMode bool
	Tmpl      tmpl.Tmpl
	*ctx.Context
	*session.Session
}
