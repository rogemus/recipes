package handlers

import (
	"errors"
	"net/http"

	"recipies.krogowski.dev/internal/consts"
	"recipies.krogowski.dev/internal/core"
	"recipies.krogowski.dev/internal/middleware"
	"recipies.krogowski.dev/internal/repository"
	"recipies.krogowski.dev/internal/validator"
)

type loginHandler struct {
	users repository.UserRepository
	requestHandler
}

func NewLoginHandler(env core.Env, userRepo repository.UserRepository) loginHandler {
	return loginHandler{
		users:          userRepo,
		requestHandler: requestHandler{Env: env},
	}
}

type loginForm struct {
	Email    string
	Password string
	validator.Validator
}

func (h *loginHandler) get(w http.ResponseWriter, r *http.Request) {
	data := h.Tmpl.NewData(r)
	data.Form = loginForm{}
	h.render(w, r, http.StatusOK, "login.tmpl", data)
}

func (h *loginHandler) post(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// bad request
		h.serverError(w, r, err)
		return
	}

	form := loginForm{
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	form.CheckField(validator.NotBlank(form.Email), "email", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", validator.FieldErr.ErrNotEmail())
	form.CheckField(validator.NotBlank(form.Password), "password", validator.FieldErr.ErrNotBlank())

	if !form.IsValid() {
		data := h.Tmpl.NewData(r)
		data.Form = form
		h.render(w, r, http.StatusUnprocessableEntity, "login.tmpl", data)
		return
	}

	id, err := h.users.Authenticate(form.Email, form.Password)

	if err != nil {
		if errors.Is(err, consts.ErrInvalidCredentials) {
			form.AddFormError(validator.FormErros.ErrInvalidCredentials())
			data := h.Tmpl.NewData(r)
			data.Form = form
			h.render(w, r, http.StatusUnprocessableEntity, "login.tmpl", data)
			return
		}

		h.serverError(w, r, err)
		return
	}

	if err = h.RenewToken(r); err != nil {
		h.serverError(w, r, err)
		return
	}

	user, err := h.users.Get(id)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	h.SetLoginUser(r, id, user.Name)

	h.SetFlashMsg(r, consts.MsgUserAuthenticeted)
	http.Redirect(w, r, "/recipies/create", http.StatusSeeOther)
}

func (h *loginHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /auth/login", midw.Dynamic.ThenFunc(h.get))
	mux.Handle("POST /auth/login", midw.Dynamic.ThenFunc(h.post))
}
