package middleware

import (
	"github.com/justinas/alice"
	"recipes.krogowski.dev/internal/core"
	"recipes.krogowski.dev/internal/repository"
)

type Midw struct {
	userRepo  repository.UserRepository
	Dynamic   alice.Chain
	Standard  alice.Chain
	Protected alice.Chain
	core.Env
}

func New(env core.Env, userRepo repository.UserRepository) *Midw {
	return &Midw{
		userRepo: userRepo,
		Env:      env,
	}
}

func (m *Midw) Init() {
	standard := alice.New(m.logRequest, m.commonHeader)
	dynamic := alice.New(m.loadAndSave, m.noSurf, m.authenticate)
	protected := dynamic.Append(m.requireAuth)

	m.Dynamic = dynamic
	m.Standard = standard
	m.Protected = protected
}
