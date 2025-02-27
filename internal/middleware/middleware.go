package middleware

import (
	"github.com/justinas/alice"
	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/repository"
)

type Midw struct {
	userRepo       repository.UserRepository
	DynamicChain   alice.Chain
	StandardChain  alice.Chain
	ProtectedChain alice.Chain
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
	dynamic := alice.New()
	protected := dynamic.Append()

	m.DynamicChain = dynamic
	m.StandardChain = standard
	m.ProtectedChain = protected
}
