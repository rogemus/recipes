package handlers

import (
	"net/http"

	"recipies.krogowski.dev/internal/consts"
	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/middleware"
)

type logoutHangler struct {
	requestHandler
}

func NewLogoutHandler(env core.Env) logoutHangler {
	return logoutHangler{
		requestHandler{Env: env},
	}
}

func (h *logoutHangler) post(w http.ResponseWriter, r *http.Request) {
	if err := h.RenewToken(r); err != nil {
		h.serverError(w, r, err)
		return
	}

	h.RemoveToken(r)
	h.SetFlashMsg(r, consts.MsgLogout)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *logoutHangler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("POST /auth/logout", midw.Protected.ThenFunc(h.post))
}
