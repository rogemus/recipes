package session

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
)

type sessionKey = string

const UserIdSessionKey = sessionKey("authenticatedUserID")
const UserNameSessionKey = sessionKey("authenticatedUserName")

type Session struct {
	Manager *scs.SessionManager
}

func New(db *sql.DB) *Session {
	manager := scs.New()
	manager.Store = sqlite3store.New(db)
	manager.Lifetime = 12 * time.Hour

	return &Session{
		Manager: manager,
	}
}

func (s *Session) SetFlashMsg(r *http.Request, msg string) {
	s.Manager.Put(r.Context(), "flash", msg)
}

func (s *Session) GetUserId(r *http.Request) int {
	id := s.Manager.GetInt(r.Context(), UserIdSessionKey)
	return id
}

func (s *Session) GetUserName(r *http.Request) string {
	name := s.Manager.GetString(r.Context(), UserNameSessionKey)
	return name
}

func (s *Session) GetFlash(r *http.Request) string {
	return s.Manager.PopString(r.Context(), "flash")
}
