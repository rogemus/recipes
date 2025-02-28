package session

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

type sessionKey = string

const UserIdSessionKey = sessionKey("authenticatedUserID")
const UserNameSessionKey = sessionKey("authenticatedUserName")

type Session struct {
	manager *scs.SessionManager
}

func New(db *sql.DB) *Session {
	manager := scs.New()
	manager.Store = postgresstore.New(db)
	manager.Lifetime = 12 * time.Hour

	return &Session{
		manager: manager,
	}
}

func (s *Session) SetFlashMsg(r *http.Request, msg string) {
	s.manager.Put(r.Context(), "flash", msg)
}

func (s *Session) GetUserId(r *http.Request) int {
	id := s.manager.GetInt(r.Context(), UserIdSessionKey)
	return id
}

func (s *Session) GetUserName(r *http.Request) string {
	name := s.manager.GetString(r.Context(), UserNameSessionKey)
	return name
}

func (s *Session) GetFlash(r *http.Request) string {
	return s.manager.PopString(r.Context(), "flash")
}

func (s *Session) SetLoginUser(r *http.Request, userId int, userName string) {
	s.manager.Put(r.Context(), UserIdSessionKey, userId)
	s.manager.Put(r.Context(), UserNameSessionKey, userName)
}

func (s *Session) RenewToken(r *http.Request) error {
	err := s.manager.RenewToken(r.Context())
	return err
}

func (s *Session) RemoveToken(r *http.Request) {
	s.manager.Remove(r.Context(), UserIdSessionKey)
}

func (s *Session) LoadAndSave(next http.Handler) http.Handler {
	return s.manager.LoadAndSave(next)
}
