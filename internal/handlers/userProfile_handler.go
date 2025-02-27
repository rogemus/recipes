package handlers

import (
	"net/http"

	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/middleware"
	"recipies.krogowski.dev/internal/repository"
)

type userProfileHandler struct {
	users repository.UserRepository
	requestHandler
}

func NewUserProfileHandler(env core.Env, userRepo repository.UserRepository) userProfileHandler {
	return userProfileHandler{
		users:          userRepo,
		requestHandler: requestHandler{Env: env},
	}
}

func (h *userProfileHandler) get(w http.ResponseWriter, r *http.Request) {
	userId := h.GetUserId(r)

	user, err := h.users.Get(userId)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := h.Tmpl.NewData(r)
	data.User = user

	h.render(w, r, http.StatusOK, "usrProfile.tmpl", data)
}

func (h *userProfileHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /usr/profile", midw.Protected.ThenFunc(h.get))
}
