package middleware

import "net/http"

func (m *Midw) loadAndSave(next http.Handler) http.Handler {
	return m.Session.Manager.LoadAndSave(next)
}
