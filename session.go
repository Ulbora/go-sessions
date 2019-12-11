package gosession

import (
	"log"
	"net/http"

	gs "github.com/gorilla/sessions"
)

//GoSession a GoSession
type GoSession struct {
	store      *gs.CookieStore
	MaxAge     int
	Secure     bool
	Name       string
	SessionKey string
}

//InitSessionStore initialize session store
func (s *GoSession) InitSessionStore() {
	if s.MaxAge == 0 {
		log.Println("using defalut max age")
		s.MaxAge = 3600 //default 3600 seconds --  1 hour
	}
	if s.Name == "" {
		log.Println("using defalut name")
		s.Name = "user-session"
	}
	if s.SessionKey == "" {
		log.Println("using defalut sesstion key")
		s.SessionKey = "554dfgdffdd11dfgf1ff1f" // default key
	}
	if s.store == nil {
		s.createSessionStore()
	}
}

// CreateSessionStore creates a sesstion
func (s *GoSession) createSessionStore() {
	s.store = gs.NewCookieStore([]byte(s.SessionKey))
	s.store.Options = &gs.Options{
		MaxAge:   s.MaxAge,
		HttpOnly: true,
		Secure:   s.Secure,
	}
}

//GetSession get session
func (s *GoSession) GetSession(r *http.Request) (*gs.Session, error) {
	return s.store.Get(r, s.Name)
}
