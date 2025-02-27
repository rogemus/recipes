package core

import (
	"log/slog"

	"recipies.krogowski.dev/internal/tmpl"
)

type Env struct {
	Logger    *slog.Logger
	DebugMode bool
	TmplCache tmpl.TmplCache
}
