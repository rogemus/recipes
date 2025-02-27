package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"

	"recipies.krogowski.dev/internal/tmpl"
)

func (h *requestHandler) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
		body   = http.StatusText(http.StatusInternalServerError)
	)

	h.Logger.Error(err.Error(), "method", method, "uri", uri)

	if h.DebugMode {
		body = fmt.Sprintf("%s\n%s", err, trace)
	}

	http.Error(w, body, http.StatusInternalServerError)
}

func (h *requestHandler) render(w http.ResponseWriter, r *http.Request, status int, page string, data tmpl.TmplData) {
	ts, ok := h.TmplCache[page]

	if !ok {
		err := fmt.Errorf("template: template %s does not exist", page)
		h.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	if err := ts.ExecuteTemplate(buf, "base", data); err != nil {
		h.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
