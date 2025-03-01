package core

import (
	"log/slog"

	"recipes.krogowski.dev/internal/ctx"
	"recipes.krogowski.dev/internal/session"
	"recipes.krogowski.dev/internal/tmpl"
)

type Env struct {
	Logger    *slog.Logger
	DebugMode bool
	Tmpl      tmpl.Tmpl
	*ctx.Context
	*session.Session
}
