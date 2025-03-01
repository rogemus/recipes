package handlers

import (
	"net/http"

	"recipes.krogowski.dev/internal/consts"
	"recipes.krogowski.dev/internal/core"
	"recipes.krogowski.dev/internal/middleware"
	"recipes.krogowski.dev/internal/repository"
	"recipes.krogowski.dev/internal/validator"
)

type signupHandler struct {
	users repository.UserRepository
	requestHandler
}

func NewSignupHandler(env core.Env, userRepo repository.UserRepository) signupHandler {
	return signupHandler{
		users:          userRepo,
		requestHandler: requestHandler{Env: env},
	}
}

type signupForm struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	validator.Validator
}

func (h *signupHandler) post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// bad request
		h.serverError(w, r, err)
		return
	}

	form := signupForm{
		Name:            r.PostForm.Get("name"),
		Email:           r.PostForm.Get("email"),
		Password:        r.PostForm.Get("password"),
		PasswordConfirm: r.PostForm.Get("passwordConfirm"),
	}

	form.CheckField(validator.NotBlank(form.Name), "name", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.NotBlank(form.Email), "email", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", validator.FieldErr.ErrNotEmail())
	form.CheckField(validator.NotBlank(form.Password), "password", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.MinChars(form.Password, 8), "password", validator.FieldErr.ErrMinLength(8))
	form.CheckField(validator.NotBlank(form.PasswordConfirm), "passwordConfirm", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.SameValue(form.Password, form.PasswordConfirm), "passwordConfirm", validator.FieldErr.ErrPassNotSame())

	if !form.IsValid() {
		data := h.Tmpl.NewData(r)
		data.Form = form
		h.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl", data)
		return
	}

	err = h.users.Insert(form.Name, form.Email, form.Password)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	h.SetFlashMsg(r, consts.MsgUserCreated)
	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func (h *signupHandler) get(w http.ResponseWriter, r *http.Request) {
	data := h.Tmpl.NewData(r)
	data.Form = signupForm{}
	h.render(w, r, http.StatusOK, "signup.tmpl", data)
}

func (h *signupHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /auth/signup", midw.Dynamic.ThenFunc(h.get))
	mux.Handle("POST /auth/signup", midw.Dynamic.ThenFunc(h.post))
}
