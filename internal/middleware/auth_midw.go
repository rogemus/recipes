package middleware

import (
	"net/http"

	"github.com/justinas/nosurf"
	"recipies.krogowski.dev/internal/ctx"
)

func (m *Midw) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.Ctx.IsAuthenticated(r) {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func (m *Midw) noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return csrfHandler
}

func (m *Midw) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := m.Session.GetUserId(r)

		if id == 0 {
			next.ServeHTTP(w, r)
			return
		}
		exists, _ := m.userRepo.Exists(id)

		// if err != nil {
		// 	// 403
		// 	m.errs.ServerError(w, r, err)
		// 	return
		// }
		//
		if exists {
			ctx := m.Ctx.CopyCtxWithKey(r, ctx.IsAuthenticatedContextKey, true)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
